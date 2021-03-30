package convert

import (
	"log"
	"net/url"
	"strings"
)

func Str2Url(urls string) *url.URL {
	u, err := url.Parse(urls)
	if err != nil {
		log.Println("[!] Str2Url Error: ", err)
		return nil
	}
	return u
}
func UrlDecode(param string) string {
	ret, err := url.QueryUnescape(param)
	if err != nil {
		log.Println("[*] UrlDecode Error: ", err)
		return ""
	}
	return ret
}
func UrlEncode(param string) string {
	ret := url.QueryEscape(param)
	return strings.ReplaceAll(ret, "+", "%20")
}

func RawEncode(param string) string {
	return strings.Replace(url.QueryEscape(param), "+", "%20", -1)
}

func RawDecode(param string) string {
	ret, err := url.QueryUnescape(strings.Replace(param, "%20", "+", -1))
	if err != nil {
		log.Println("RawDecode Error: ", err)
		return ""
	}
	return ret
}
