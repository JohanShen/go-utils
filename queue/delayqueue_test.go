package queue

import (
	"fmt"
	"testing"
)

func TestSlic(t *testing.T){
	ary1 := make([]int, 0, 100)

	t.Logf("%d",len(ary1))
	t.Logf("%p",ary1)
	ary2 := append(ary1, 0)
	for i:=0; i<150;i++ {
		ary2 = append(ary2, i)
		//ary2[i] = i
		t.Logf("%p", ary2)
		t.Log(len(ary2))
	}

	ary3 := ary2[:10]
	t.Logf("ary3 : %p", ary3)

	ary4 := ary2[10:]
	t.Logf("ary4 : %p",ary4)


}


func callback(datas []interface{}) error{
	for i:=0;i<len(datas);i++{
		fmt.Print(datas[i])
	}
	return nil
}

func TestDelayQueue_Push(t *testing.T) {

	 q  := NewDelayQueue(150, callback)
	go func(){
		for i:=0;i<100;i++{
			q.Push(i)
		}
	}()
}

