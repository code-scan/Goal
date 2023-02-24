package Ghttp

import (
	"log"
)

func (h *Http) SetHeader(key string, value string) *Http {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	if h.HttpRequest == nil {
		log.Println("[!]HttpRequest Is not Init")
		return h
	}
	if head := h.HttpRequest.Header.Get(key); head != "" {
		h.HttpRequest.Header.Set(key, value)
	} else {
		h.HttpRequest.Header.Add(key, value)
	}
	return h
}
func (h *Http) SetUserAgent(agent string) {
	h.SetHeader("User-Agent", agent)
}

//设置请求content type
func (h *Http) SetContentType(s string) *Http {
	switch s {
	case "form":
		h.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	case "json":
		h.SetHeader("Content-Type", "application/json;charset=UTF-8")
	default:
		h.SetHeader("Content-Type", s)
	}
	return h
}
