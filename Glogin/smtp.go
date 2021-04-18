package Glogin

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

func SmtpLogin(host, port, email, password string, ssl bool) bool {
	hostname := fmt.Sprintf("%s:%s", host, port)
	domain := strings.Split(email, "@")[1]
	auth := smtp.PlainAuth("", email, password, domain)
	var smtpClient *smtp.Client
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
			defer conn.Close()
			smtpClient, err = smtp.NewClient(conn, domain)
		}

	} else {

		if tcpConn, err := net.Dial("tcp", hostname); err != nil {
			log.Println("net.Dial: ", err)
			return false
		} else {
			defer tcpConn.Close()
			smtpClient, err = smtp.NewClient(tcpConn, domain)
		}
	}

	if err != nil {
		log.Println("smtp.NewClient: ", err)
		return false
	}
	defer smtpClient.Close()
	err = smtpClient.Auth(auth)
	if err != nil {
		log.Println("Auth: ", err)
		return false
	}
	return true
}
