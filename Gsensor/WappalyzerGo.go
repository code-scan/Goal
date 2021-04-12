package Gsensor

import (
	"fmt"
	"github.com/code-scan/Goal/Ghttp"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type WappalyzerGo struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	http     Ghttp.Http
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
	resp, err := http.DefaultClient.Get(s.Domain)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	for k := range fingerprints {
		r[k] = ""
	}
	return r
}

func (s *WappalyzerGo) Login(_ bool) bool {
	return true
}
