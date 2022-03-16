package utils

import (
	"math/rand"
	"net"
	"os"
	"time"
)

// 当前时间
func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 当前天
func Day() string {
	return time.Now().Format("2006-01-02")
}

// 微妙
func MicroSecond() int64 {
	return time.Now().UnixNano() / 1000000
}

func HostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknow"
	}
	return hostname
}

func Ip() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

// 随机数
func Random(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
