package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
)

func GetMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	if len(netInterfaces) == 0 {
		return "", errors.New("无法获取到正确的MAC地址，网卡数量为空。")
	}

	str := make([]string, len(netInterfaces))
	for i := 0; i < len(netInterfaces); i++ {
		//fmt.Println(netInterfaces[i])
		if (netInterfaces[i].Flags&net.FlagUp) != 0 && (netInterfaces[i].Flags&net.FlagLoopback) == 0 {
			adds, _ := netInterfaces[i].Addrs()
			for _, address := range adds {
				inet, ok := address.(*net.IPNet)
				//fmt.Println(inet.IP)
				if ok && inet.IP.IsGlobalUnicast() {
					// 如果IP是全局单拨地址，则返回MAC地址
					mac := netInterfaces[i].HardwareAddr.String()
					str[i] = mac
				}
			}
		}
	}

	return strings.Join(str, "|"), nil
}

func GetMachineId() string {
	mac, err := GetMACAddress()
	if err != nil {
		mac = ""
	}
	name, err := os.Hostname()
	if err != nil {
		name = "default"
	}
	pid := os.Getpid()
	key := fmt.Sprintf("%v|%v", mac, name)
	val := uint64(HashCode(key)) + uint64(pid)

	if v, err := Oct2Any(val, 62); err != nil {
		return ""
	} else {
		return v
	}
}
