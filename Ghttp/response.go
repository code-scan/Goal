package Ghttp

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//发送请求
func (h *Http) Execute() *http.Response {
	var err error
	h.HttpClient.Transport = &h.HttpTransport
	h.HttpResponse, err = h.HttpClient.Do(h.HttpRequest)
	if err != nil {
		log.Println("[!] Http Execute Error : ", err)
		h.HttpResponse = nil
		return nil
	}
	return h.HttpResponse
}

// string的返回值
func (h *Http) Text() (string, error) {
	var result []byte
	var err error
	if h.HttpResponse == nil {
		return "", err
	}
	result, err = ioutil.ReadAll(h.HttpResponse.Body)
	if err != nil {
		log.Println("[!]Text Error: ", err)
		return "", err
	}
	defer h.HttpResponse.Body.Close()
	return string(result), err
}

// byte的返回值
func (h *Http) Byte() ([]byte, error) {
	var result []byte
	var err error
	if h.HttpResponse == nil {
		return result, err
	}
	result, err = ioutil.ReadAll(h.HttpResponse.Body)
	if err != nil {
		log.Println("[!]Text Error: ", err)
		return result, err
	}
	defer h.HttpResponse.Body.Close()
	return result, err
}
func (h *Http) SaveToFile(file string) (bool, error) {
	var err error
	if h.HttpResponse == nil {
		log.Println("[!] HttpResponse Is Closed")
		return false, err
	}
	var f *os.File
	f, err = os.Create(file)
	if h.HttpResponse.Body != nil && err == nil {
		defer h.HttpResponse.Body.Close()
		defer f.Close()
		_, err = io.Copy(f, h.HttpResponse.Body)
	}
	if err == nil {
		return true, err
	}
	return false, err

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
