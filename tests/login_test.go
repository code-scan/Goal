package test

import (
	"log"
	"testing"

	"github.com/code-scan/Goal/Glogin"
)

func TestLogin(t *testing.T) {
	r := Glogin.FtpLogin("127.0.0.1", "2121", "anonymous", "anonymous")
	log.Println(r)
}
func TestSmtp(t *testing.T) {
	r := Glogin.SmtpLogin("smtp.126.com", "465", "123123", "123123", true)
	log.Println(r)
}
func TestPop3(t *testing.T) {
	r := Glogin.Pop3Login("pop.126.com", "995", "123123", "123123", true)
	log.Println(r)
}
func TestSQL(t *testing.T) {
	r := Glogin.SqlQuery("127.0.0.1", "3306", "root", "123123", "mysql", "select * from users limit 10")
	log.Println(r)

}
