package Gsensor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/code-scan/Goal/Ghttp"
	"log"
	"strings"
)

type Fofa struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	result   Result
	http     Ghttp.Http
}

type FofaResult struct {
	Error   bool       `json:"error"`
	Mode    string     `json:"mode"`
	Page    int        `json:"page"`
	Query   string     `json:"query"`
	Results [][]string `json:"results"`
	Size    int        `json:"size"`
}

func (s *Fofa) GetInfo() string {
	return "FofaImpl ver 0.1 with  " + s.Type
}
func (s *Fofa) SetType(Type string) {
	s.Type = Type
}
func (s *Fofa) SetDomain(domain string) {
	s.Domain = domain
}
func (s *Fofa) SetAccount(account string) {
	s.UserName = account
}
func (s *Fofa) SetPassword(password string) {
	s.PassWord = password
}
func (s *Fofa) GetResult() Result {
	switch s.Type {
	case "subdomain":
		s.GetSubDomain(1)
	case "sameserver":
		s.GetIp(1)
	case "ports":
		s.GetPorts()
	}
	return s.result
}
func (s *Fofa) Login(ReLogin bool) bool {
	//s.http = HttpHelper{}
	return true
}

func (s *Fofa) GetSubDomain(i int) {
	s.result = Result{}
	resp := s.send(`domain="` + s.Domain + `"`)
	fofaresult := FofaResult{}
	json.Unmarshal(resp, &fofaresult)
	for _, v := range fofaresult.Results {
		host := v[0]
		if strings.Contains(host, "://") {
			host = strings.Split(host, "://")[1]
		}
		s.result[host] = v[1]
	}

}

func (s *Fofa) GetIp(i int) {
	s.result = Result{}
	resp := s.send(`ip="` + s.Domain + `" && type=subdomain`)
	fofaresult := FofaResult{}
	json.Unmarshal(resp, &fofaresult)
	for _, v := range fofaresult.Results {
		host := v[0]
		if strings.Contains(host, "://") {
			host = strings.Split(host, "://")[1]
		}
		s.result[host] = v[1]
	}
}

func (s *Fofa) GetPorts() {
	s.result = Result{}
	resp := s.send(`ip="` + s.Domain + `" && type=service`)
	fofaresult := FofaResult{}
	json.Unmarshal(resp, &fofaresult)
	for _, v := range fofaresult.Results {
		host := v[0]
		if strings.Contains(host, "://") {
			host = strings.Split(host, "://")[1]
		}
		s.result[host] = v[1]
	}
}
func (s *Fofa) send(query string) []byte {
	var result []byte
	query = base64.StdEncoding.EncodeToString([]byte(query))
	urls := fmt.Sprintf("https://fofa.so/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=10000", s.UserName, s.PassWord, query)
	s.http.New("GET", urls)
	s.http.Execute()
	if ret := s.http.Execute(); ret == nil {
		return result
	}
	resp, err := s.http.Byte()
	if err != nil {
		log.Println(err)
		return result
	}
	return resp
}
