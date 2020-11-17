package utils

import (
	"fmt"
	"sync/atomic"
	"time"
)

type idGenerator struct {
	num       uint64
	ch        chan string
	timestamp uint64
	mid       string
}

var IdGenerator = NewIdGenerator()

func NewIdGenerator() *idGenerator {
	g := &idGenerator{
		num:       1,
		ch:        make(chan string),
		mid:       GetMachineId(),
		timestamp: makeTimestamp(),
	}
	go g.make()
	return g
}

func makeTimestamp() uint64 {
	return uint64(time.Now().UnixNano() / 1e3)
}

func (g *idGenerator) make() {
	for {
		oct := g.timestamp*1000 + g.num
		if id, err := Oct2Any(oct, 62); err == nil {
			//fmt.Println(oct, g.timestamp, g.num, id)
			atomic.AddUint64(&g.num, 1)
			if g.num > 9999 {
				g.timestamp = makeTimestamp()
				atomic.AddUint64(&g.num, -g.num)
			}
			//atomic.StoreUint64(&g.num, 0)

			g.ch <- fmt.Sprintf("%v%v", g.mid, id)
		}
	}
}

func (g *idGenerator) Next(num int) []string {
	r := make([]string, num)
	i := 0
	for v := range g.ch {
		r[i] = v
		i++
		if i == num {
			break
		}
	}
	return r
}

func (g *idGenerator) NextId() string {
	return <-g.ch
}
