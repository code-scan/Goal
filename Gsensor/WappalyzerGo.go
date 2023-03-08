package Gsensor

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"git.dev.me/jerry/Goal/Ghttp"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

type WappalyzerGo struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	Http     Ghttp.Http
}

func (s *WappalyzerGo) GetInfo() string {
	return "WappalyzerGo ver 0.1 with  " + s.Type
}

func (s *WappalyzerGo) SetDomain(domain string) {
	if strings.Contains(domain, "http") == false {
		domain = fmt.Sprintf("http://%s", domain)
	}
	s.Domain = domain
}

func (s *WappalyzerGo) SetAccount(username string) {
	s.UserName = username
}

func (s *WappalyzerGo) SetPassword(password string) {
	s.PassWord = password
}

func (s *WappalyzerGo) SetType(type_ string) {
	s.Type = type_
}

func (s *WappalyzerGo) GetResult() Result {
	r := Result{}
	if s.Type != "finger" {
		return r
	}
	//resp, err := http.DefaultClient.Get(s.Domain)
	s.Http.New("GET", s.Domain)
	resp := s.Http.Execute()
	defer s.Http.Close()
	if resp == nil {
		log.Println("resp nil")
		return r
	}
	data, err := ioutil.ReadAll(resp.HttpResponse.Body) // Ignoring error for example
	if err != nil {
		log.Println(err)
		return r
	}
	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		log.Println(err)
		return r
	}
	fingerprints := wappalyzerClient.Fingerprint(resp.HttpResponse.Header, data)
	for k := range fingerprints {
		r[k] = ""
	}
	return r
}

func (s *WappalyzerGo) Login(_ bool) bool {
	return true
}
