package Glogin

import (
	"crypto/tls"
	"fmt"
	"github.com/taknb2nch/go-pop3"
	"log"
	"net"
)

func Pop3Login(host, port, email, password string, ssl bool) bool {

	hostname := fmt.Sprintf("%s:%s", host, port)
	//domain := strings.Split(email, "@")[1]
	var pop3Client *pop3.Client
	var err error

	if ssl {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}
		if conn, err := tls.Dial("tcp", hostname, tlsconfig); err != nil {
			log.Println("tls.Dial: ", err)
			return false
		} else {
			pop3Client, err = pop3.NewClient(conn)
		}

	} else {

		if tcpConn, err := net.Dial("tcp", hostname); err != nil {
			log.Println("net.Dial: ", err)
			return false
		} else {
			pop3Client, err = pop3.NewClient(tcpConn)
		}
	}
	//log.Println(err)
	//log.Println(pop3Client)
	//return false
	if err != nil || pop3Client == nil {
		log.Println("pop3Client: ", err)
		return false
	}
	if err := pop3Client.User(email); err != nil {
		log.Println("User: ", err)
		return false
	}
	if err := pop3Client.Pass(password); err != nil {
		log.Println("Pass: ", err)
		return false
	}
	return true
}
