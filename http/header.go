package http

import (
	"log"
)

func (h *Http) SetHeader(key string, value string) {
	if h.HttpRequest == nil {
		log.Panicln("HttpRequest Is not Init")
		return
	}
	if head := h.HttpRequest.Header.Get(key); head != "" {
		h.HttpRequest.Header.Set(key, value)
	} else {
		h.HttpRequest.Header.Add(key, value)
	}
}
func (h *Http) SetUserAgent(agent string) {
	h.SetHeader("User-Agent", agent)
}


//设置请求content type
func (h *Http) SetContentType(s string) {
	switch s {
	case "form":
		h.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	case "json":
		h.SetHeader("Content-Type", "application/json;charset=UTF-8")
	default:
		h.SetHeader("Content-Type", s)
	}
}
