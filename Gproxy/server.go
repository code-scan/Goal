package Gproxy

import (
	"fmt"
	"github.com/hashicorp/yamux"
	"io"
	"log"
	"net"
	"os"
)

var session *yamux.Session

func ClientWait(port string) {
	hostClient := "0.0.0.0"
	portClient := port
	client, err := net.Listen("tcp", fmt.Sprintf("%s:%s", hostClient, portClient))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
	fmt.Println("client list on :7777")
	for {
		conn, err := client.Accept()
		if err != nil {
			fmt.Println("clien error")
			fmt.Println(err)
			continue
		}
		session, err = yamux.Server(conn, nil)
	}
}
func ServerWait(port string) {
	hostServer := "0.0.0.0"
	portServer := port
	server, err := net.Listen("tcp", fmt.Sprintf("%s:%s", hostServer, portServer))
	if err != nil {
		fmt.Println(err, err.Error())
		os.Exit(0)
	}
	fmt.Println("server list on :8888")
	for {
		serverConn, err := server.Accept()
		if err != nil {
			fmt.Println("server error")
			fmt.Println(err)
		}
		if session == nil {
			serverConn.Close()
			continue
		}
		stream, err := session.Open()
		if err != nil {
			fmt.Println("session error")
			fmt.Println(err)
			break
		}
		go func() {
			log.Println("Starting to copy serverConn to stream")
			io.Copy(serverConn, stream)
			serverConn.Close()
		}()
		go func() {
			log.Println("Starting to copy stream to serverConn")
			io.Copy(stream, serverConn)
			serverConn.Close()
		}()

	}
}

//go clientWait(os.Args[1])
//serverWait(os.Args[2])
