package Ghttp

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Http struct {
	HttpClient      http.Client    // http客户端
	HttpRequest     *http.Request  // Ghttp 请求
	HttpResponse    *http.Response // Ghttp 返回值
	HttpRequestUrl  string         //请求的url
	HttpRequestType string         // 请求方法，GET/POST
	HttpContentType string         // 请求类型 json/form-url-encoide
	HttpBody        io.Reader      // 返回内容
	HttpTransport   http.Transport
	Cookie          *cookiejar.Jar //cookie的值
	isSession       bool           //是否创建session
}

//var HttpClient Http
//
//func init() {
//	HttpClient = Http{}
//}

// 新建一个请求
func (h *Http) New(method, urls string) error {
	var err error
	h.HttpRequestUrl = urls
	h.HttpRequestType = method
	// 初始化http client 如果开启了session，则传入cookie jar
	if h.isSession {
		if h.Cookie == nil {
			h.Cookie, _ = cookiejar.New(nil)
		}
		h.HttpClient.Jar = h.Cookie
	}
	h.HttpRequest, err = http.NewRequest(h.HttpRequestType, h.HttpRequestUrl, h.HttpBody)
	return err

	//if h.HttpRequest == nil {
	//	h.HttpRequest, err = Ghttp.NewRequest(h.HttpRequestType, h.HttpRequestUrl, h.HttpBody)
	//} else if h.isSession { //如果不是第一次请求 并且开启了session 则复用之前的request即可
	//	var uri *url.URL
	//	uri, err = url.Parse(h.HttpRequestUrl)
	//	if err != nil {
	//		return err
	//	}
	//	h.HttpRequest.Method = h.HttpRequestType
	//	h.HttpRequest.URL = uri
	//}

}
func (h *Http) SetTimeOut(t int) {
	td := time.Duration(t)
	h.HttpTransport.TLSHandshakeTimeout = td * time.Second
	h.HttpTransport.ResponseHeaderTimeout = td * time.Second
	h.HttpTransport.IdleConnTimeout = td * time.Second
	h.HttpTransport.ExpectContinueTimeout = td * time.Second

}
func (h *Http) IgnoreSSL() {
	h.HttpTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
}
