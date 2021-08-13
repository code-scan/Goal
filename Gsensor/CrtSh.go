package Gsensor

import (
	"github.com/code-scan/Goal/Ghttp"
	"log"
	"regexp"
	"strings"
)

type CrtSh struct {
	Domain string
	Type   string
	result Result
	http   Ghttp.Http
}

func (s *CrtSh) GetInfo() string {
	return "CrtSh ver 0.1 with  " + s.Type

}

func (s *CrtSh) SetDomain(domain string) {
	s.Domain = domain
}

func (s *CrtSh) SetAccount(_ string) {
	return
}

func (s *CrtSh) SetPassword(_ string) {
	return
}

func (s *CrtSh) SetType(type_ string) {
	s.Type = type_
}

func (s *CrtSh) GetResult() Result {
	s.result = Result{}
	if s.Type != "subdomain" {
		return s.result
	}

	s.http.Get("https://crt.sh/?q=" + s.Domain)
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Text()
	if err != nil {
		log.Println("[!] GetResult Error: ", err)
		return s.result
	}
	reg := regexp.MustCompile(`<TD>(.*?)</TD>`)
	r := reg.FindAllString(ret, -1)
	for _, v := range r {
		if strings.Contains(v, s.Domain) {
			v = strings.ReplaceAll(v, "<TD>", "")
			v = strings.ReplaceAll(v, "</TD>", "")
			vs := strings.Split(v, "<BR>")
			for _, d := range vs {
				d = strings.ReplaceAll(d, "*.", "")
				s.result[d] = "未解析"
			}
		}
	}

	return s.result
}

func (s *CrtSh) Login(_ bool) bool {
	return true
}
