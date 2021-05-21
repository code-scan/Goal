package Ghttp

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

//发送请求
func (h *Http) Execute() *http.Response {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	//fmt.Printf("HttpTransport: %p \n", h.HttpTransport)
	var err error
	h.HttpClient.Transport = h.HttpTransport
	h.HttpResponse, err = h.HttpClient.Do(h.HttpRequest)
	if err != nil {
		log.Println("[!] Http Execute Error : ", err)
		h.HttpResponse = nil
		return nil
	}
	return h.HttpResponse
}

// 关闭请求与body
func (h *Http) Close() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if h.HttpResponse != nil {
		h.HttpResponse.Body.Close()
		//log.Println("CLose(): ", h.HttpResponse.Body)
		h.readAll()
	}
	if h.CtxCancel != nil {
		h.CtxCancel()
	}

}

//获取返回头
func (h *Http) GetRespHead(key string) string {
	if h.HttpResponse != nil {
		return h.HttpResponse.Header.Get(key)
	}
	return ""
}

func (h *Http) readAll() ([]byte, error) {
	buffer := h.Pool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer func() {
		if buffer != nil {
			h.Pool.Put(buffer)
			buffer = nil
		}
	}()
	_, err := io.Copy(buffer, h.HttpResponse.Body)
	if err != nil && err != io.EOF {
		log.Printf("adapter io.copy failure error:%v \n", err)
		return nil, fmt.Errorf("adapter io.copy failure error:%v", err)
	}
	defer h.HttpResponse.Body.Close()
	return buffer.Bytes(), nil
}

// string的返回值
func (h *Http) Text() (string, error) {
	var result []byte
	var err error
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in Text : ", r)
		}
	}()
	if h.HttpResponse == nil {
		return "", err
	}
	if h.HttpResponse.Body == nil {
		return "", err
	}
	if result, err = h.readAll(); err != nil {
		log.Println(err)
	}
	return string(result), err
}

// byte的返回值
func (h *Http) Byte() ([]byte, error) {
	var result []byte
	var err error
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in Byte : ", r)
		}
	}()
	if h.HttpResponse == nil {
		return result, err
	}
	if h.HttpResponse.Body == nil {
		return result, err
	}
	if result, err = h.readAll(); err != nil {
		log.Println(err)
	}
	return result, err
}
func (h *Http) SaveToFile(file string) (bool, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in SaveToFile : ", r)
		}
	}()
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
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if h.HttpResponse != nil {
		return h.HttpResponse.StatusCode
	}
	return -1
}
func (h *Http) RespCookie() string {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in RespCookie : ", r)
		}
	}()
	if h.HttpResponse == nil || h.HttpResponse.Header == nil {
		return ""
	}
	cookies := h.HttpResponse.Header.Values("set-cookie")
	cks := strings.Join(cookies, "; ")
	return cks
}
