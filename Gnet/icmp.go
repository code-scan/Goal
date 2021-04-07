package Gnet

import "net"

func PingHost(host string) bool {
	raddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return false
	}
	conn, err := net.DialIP("ip4:icmp", nil, raddr)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
