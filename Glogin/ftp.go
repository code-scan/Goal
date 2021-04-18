package Glogin

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func FtpLogin(host, port, username, password string) bool {
	connStr := fmt.Sprintf("%s:%s", host, port)
	c, err := ftp.Dial(connStr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return false
	}

	err = c.Login(username, password)
	if err != nil {
		return false
	}
	defer c.Quit()
	return true
}
