package Gnet

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func TcpPortStatus(ip string, port int, timeout int) bool {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func GetIPList(ip string) []string {
	if strings.Contains(ip, "/") == false {
		return []string{
			ip,
		}
	}
	var ipList []string
	ipdata, ipAddress, err := net.ParseCIDR(ip)
	if err != nil {
		return ipList
	}
	for IP := ipdata.Mask(ipAddress.Mask); ipAddress.Contains(IP); inc(IP) {
		ipList = append(ipList, IP.String())
	}
	return ipList
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}



