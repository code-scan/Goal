package Gproxy

import (
	"fmt"
	"log"
	"net"
)
import (
	"github.com/armon/go-socks5"
	"github.com/hashicorp/yamux"
)
var session_client *yamux.Session

func RunProxy(remoteAddr string){
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}
	fmt.Println("forward_conn")
	forward_conn,err:=net.Dial("tcp",remoteAddr)
	if err!=nil{
		return
	}
	fmt.Println("forward_success")
	session_client, err = yamux.Server(forward_conn, nil)
	for{
		stream, err := session_client.Accept()
		if err!=nil{
			fmt.Println("session_client err")
			fmt.Println(err)
			return
		}
		log.Println("conn to socks5")
		go func() {
			err = server.ServeConn(stream)
			if err != nil {
				log.Println(err)
			}
		}()
	}

}