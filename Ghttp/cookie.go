package Ghttp

import (
	"net/http"
	"strings"
)

func (h *Http) AddCookieJar(key, val string) {
	ck := http.Cookie{
		Name:   key,
		Value:  val,
		Domain: "." + h.HttpRequest.Host,
		Path:   "/",
	}
	if h.isSession{
		h.HttpClient.Jar.SetCookies(h.HttpRequest.URL, []*http.Cookie{&ck})
	}else{
		h.Cookie.SetCookies(h.HttpRequest.URL, []*http.Cookie{&ck})
	}


}
func (h *Http) SetCookie(cookie string) {
	if h.Cookie != nil {
		cookies := strings.Split(cookie, ";")
		for _, v := range cookies {
			kv := strings.Split(strings.TrimSpace(v), "=")
			if len(kv) != 2 {
				continue
			}
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			h.AddCookieJar(key, value)

		}
	} else {
		h.SetHeader("cookie", cookie)

	}
}
