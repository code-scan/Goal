package test

import (
	"github.com/code-scan/Goal/Ghttp"
	"log"
	"net/url"
	"os"
	"runtime/pprof"
	"testing"
	"time"
)

func init() {
	log.Println("run test")
}
func TestHttp(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Ghttp test get")
	httpClient := Ghttp.Http{}
	//httpClient.SetTimeOut(1)
	httpClient.Session()
	//httpClient.SetSocksProxy("127.0.0.1:6153")
	//httpClient.SetProxy("http://127.0.0.1:6152")
	//httpClient.SetProxy("socks5://ss:ss@127.0.0.1:6153")
	//
	//httpClient.Get("https://httpbin.org/get")
	//httpClient.SetPostString("1123=123123")
	//httpClient.SetUserAgent("123123")
	//httpClient.SetCookie("1=asdasdasd")
	//httpClient.Execute()
	//log.Println(httpClient.Text())

	httpClient.Get("https://www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png")
	httpClient.Byte()
	httpClient.SaveToFile("tmp.png")
	//log.Println(httpClient.Text())

}
func TestPst(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("Ghttp test post")
	httpClient := Ghttp.New()
	httpClient.Session()
	p := make(map[string]interface{})
	p["1234"] = "1234"
	p2 := make(map[string]interface{})
	p2["key2"] = "value2"
	p["2"] = p2
	log.Println(p)
	p3 := url.Values{
		"name": {"values"},
	}
	httpClient.Post("https://baidu.com", p3)
	httpClient.DontRedirect()
	httpClient.SetCookie("a=123123;123213=cccc;")
	//httpClient.Close()
	httpClient.Execute()
	defer httpClient.Close()
	httpClient.Close()

	httpClient.Close()

	log.Println(httpClient.Text())
	log.Println(httpClient.StatusCode())
	log.Println(httpClient.RespCookie())
	//httpClient.HttpResponse.Body.Close()
	//httpClient.HttpResponse.Body.Close()
	//httpClient.HttpResponse.Body.Close()
	//httpClient.Text()
	//httpClient.Close()
	//httpClient.Close()
	log.Println(httpClient.GetRespHead("cookie"))
	log.Println("log post")
}
func TestHead(t *testing.T) {
	log.Println("Ghttp test head")
	httpClient := Ghttp.Http{}
	go func() {
		httpClient.Head("https://httpbin.org/head")
		httpClient.Execute()
		log.Println(httpClient.StatusCode())

	}()
	go func() {
		httpClient.Head("https://httpbin.org/get")
		httpClient.Execute()
		log.Println(httpClient.StatusCode())

	}()
	time.Sleep(time.Duration(time.Second * 5))
	log.Println(httpClient.StatusCode())
}
func TestManyReq(t *testing.T) {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	var task = make(chan string, 100)
	for i := 0; i < 100; i++ {
		go func() {
			c := Ghttp.New()
			for {
				tt := <-task
				log.Println("GET Task : ", tt)
				c.New("GET", tt)
				c.DontRedirect()
				c.Execute()
				location := c.GetRespHead("location")
				log.Println(location)
				c.Text()
				defer c.Close()
			}
		}()
	}
	for i := 0; i < 1000; i++ {
		task <- "http://qq.com"
	}
}
