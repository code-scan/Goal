package test

import (
	"github.com/code-scan/Goal/Ghttp"
	"log"
	"testing"
)

func TestCookie(t *testing.T){
	var http = Ghttp.Http{}
	http.Session()
	http.New("GET","https://www.ti.com")
	http.SetProxy("http://127.0.0.1:8080")
	http.SetCookie("a=1")
	http.Execute()
	log.Println(http.HttpClient.Jar)


	http.New("GET","https://www.ti.com")
	http.SetCookie("a=1234")
	http.Execute()
	log.Println(http.HttpClient.Jar)

	http.New("GET","https://www.ti.com")
	http.SetCookie("abb=1")
	http.Execute()
	log.Println(http.HttpClient.Jar)


	http.New("GET","https://www.ti.com")
	http.SetCookie("cca=ccc1")
	http.Execute()
	log.Println(http.HttpClient.Jar)

}
