package test

import (
	"log"
	"testing"

	"github.com/code-scan/Goal/Gnet"
)

func TestTcp(t *testing.T) {
	log.Println(111)
	r := Gnet.GetIPList("192.168.1.1/24")
	log.Println(r)
	rr := Gnet.TcpPortStatus("127.0.0.1", 80, 30)
	log.Println(rr)
}
func TestICMP(t *testing.T) {
	rr := Gnet.PingHost("127.0.0.1")
	log.Println(rr)
}
