package Ghttp

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

type Http struct {
	HttpClient      *http.Client   // http客户端
	HttpRequest     *http.Request  // Ghttp 请求
	HttpResponse    *http.Response // Ghttp 返回值
	HttpRequestUrl  string         //请求的url
	HttpRequestType string         // 请求方法，GET/POST
	HttpContentType string         // 请求类型 json/form-url-encoide
	HttpBody        io.Reader      // 返回内容
	HttpTransport   *http.Transport
	Cookie          *cookiejar.Jar //cookie的值
	isSession       bool           //是否创建session
	Ctx             context.Context
	CtxCancel       context.CancelFunc
	//Pool            *sync.Pool
	err error
}
type Headers map[string]string

//var HttpClient Http
//
//func init() {
//	HttpClient = Http{}
//}
var transport http.Transport
var pool sync.Pool

func init() {

	transport = http.Transport{

		//DisableKeepAlives: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	pool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 4096))
		},
	}

	//log.Println("init http")
}
func New() *Http {
	c := Http{}
	c.HttpTransport = &transport
	return &c
}

// 新建一个请求
func (h *Http) New(method, urls string) *Http {
	var err error
	h.err = nil
	h.Ctx, h.CtxCancel = context.WithCancel(context.Background())
	h.HttpRequestUrl = urls
	h.HttpRequestType = method
	// 初始化http client 如果开启了session，则传入cookie jar
	if h.HttpTransport == nil {
		//log.Println("new transport")
		h.HttpTransport = &transport
	}
	h.HttpClient = &http.Client{Transport: h.HttpTransport}

	if h.isSession {
		if h.Cookie == nil {
			h.Cookie, _ = cookiejar.New(nil)
		}
		if h.HttpClient.Jar == nil {
			h.HttpClient.Jar = h.Cookie
		}
	}

	h.SetTimeOut(30)
	h.IgnoreSSL()
	h.HttpRequest, err = http.NewRequest(h.HttpRequestType, h.HttpRequestUrl, h.HttpBody)
	h.HttpRequest.WithContext(h.Ctx)
	h.err = err
	return h
	// return err
}
func (h *Http) SetTimeOut(t int) {
	td := time.Duration(t)
	h.HttpTransport.TLSHandshakeTimeout = td * time.Second
	h.HttpTransport.ResponseHeaderTimeout = td * time.Second
	h.HttpTransport.IdleConnTimeout = td * time.Second
	h.HttpTransport.ExpectContinueTimeout = td * time.Second
	//h.HttpTransport.DisableKeepAlives = true
}
func (h *Http) IgnoreSSL() {
	h.HttpTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
}

func (h *Http) SetProxy(proxyUrl string) *Http {
	u, _ := url.Parse(proxyUrl)
	switch u.Scheme {
	case "https":
		//log.Println("use proxy", u.Scheme)
		h.HttpTransport.Proxy = http.ProxyURL(u)
	case "http":
		//log.Println("use proxy", u.Scheme)
		h.HttpTransport.Proxy = http.ProxyURL(u)
	case "socks5":
		pwd, _ := u.User.Password()
		auth := proxy.Auth{
			User:     u.User.Username(),
			Password: pwd,
		}

		h.SetSocksProxy(u.Host, &auth)

	}
	return h

}
func (h *Http) SetSocksProxy(proxyUrl string, auth *proxy.Auth) *Http {
	baseDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	dialSocksProxy, err := proxy.SOCKS5("tcp", proxyUrl, auth, baseDialer)
	if err != nil {
		log.Println("[!]SetSocksProxy Error: ", err)
		h.err = err
		return h
	}
	if contextDialer, ok := dialSocksProxy.(proxy.ContextDialer); ok {
		h.HttpTransport.DialContext = contextDialer.DialContext
	}
	return h
}

func (h *Http) DontRedirect() *Http {
	h.HttpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return h
}
func (h *Http) DontKeepAlive() *Http {
	h.HttpTransport.DisableKeepAlives = true
	return h
}

func (h *Http) Error() error {
	return h.err
}

func Get(url, proxy string, headers Headers) ([]byte, error) {
	http := New()
	http.Get(url)
	for k, v := range headers {
		http.SetHeader(k, v)
	}
	if proxy != "" {
		http.SetProxy(proxy)
	}
	http.IgnoreSSL()
	http.SetTimeOut(30)
	resp := http.Execute()
	if resp == nil {
		return nil, fmt.Errorf("http connect error")
	}
	if http.HttpResponse.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(http.HttpResponse.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return ioutil.ReadAll(reader)
	}
	return http.Byte()

}
func Post(url string, data interface{}, proxy string, headers Headers) ([]byte, error) {
	http := New()
	http.Post(url, data)
	for k, v := range headers {
		http.SetHeader(k, v)
	}
	if proxy != "" {
		http.SetProxy(proxy)
	}
	http.IgnoreSSL()
	if strings.HasPrefix(data.(string), "{") == false {
		http.SetContentType("form")
	}
	http.Execute()
	if http.HttpResponse.Header.Get("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(http.HttpResponse.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		return ioutil.ReadAll(reader)
	}
	return http.Byte()
}
