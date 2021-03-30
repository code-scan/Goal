package Ghttp

import (
	"io/ioutil"
	"net/http"
	"strings"
)

//发送请求
func (h *Http) Execute() *http.Response {
	var err error
	h.HttpClient.Transport = &h.HttpTransport
	h.HttpResponse, err = h.HttpClient.Do(h.HttpRequest)
	if err != nil {
		return &http.Response{}
	}
	return h.HttpResponse
}

// string的返回值
func (h *Http) Text() (string, error) {
	var result []byte
	var err error
	if h.HttpResponse != nil {
		result, err = ioutil.ReadAll(h.HttpResponse.Body)
	}
	return string(result), err
}

// byte的返回值
func (h *Http) Byte() ([]byte, error) {
	var result []byte
	var err error
	if h.HttpResponse != nil {
		result, err = ioutil.ReadAll(h.HttpResponse.Body)
	}
	return result, err
}

// statuscode
func (h Http) StatusCode() int {
	if h.HttpResponse != nil {
		return h.HttpResponse.StatusCode
	}
	return -1
}
func (h *Http) RespCookie() string {
	cookies := h.HttpResponse.Header.Values("set-cookie")
	cks := strings.Join(cookies, "; ")
	return cks
}
