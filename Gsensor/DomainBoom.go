package Gsensor

import (
	"encoding/json"
	"fmt"
	"log"

	"git.dev.me/jerry/Goal/Ghttp"
)

type DomainBoom struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	result   Result
	http     Ghttp.Http
}
type DomainBoomResult []string

func (s *DomainBoom) GetInfo() string {
	return "DomainBoom ver 0.1 with  " + s.Type
}

func (s *DomainBoom) SetDomain(domain string) {
	s.Domain = domain
}

func (s *DomainBoom) SetAccount(_ string) {
	return
}

func (s *DomainBoom) SetPassword(password string) {
	s.PassWord = password
}

func (s *DomainBoom) SetType(type_ string) {
	s.Type = type_
}

func (s *DomainBoom) GetResult() Result {
	s.result = Result{}
	if s.Type != "subdomain" {
		return s.result
	}
	uri := fmt.Sprintf("http://domain.f5.pm/api.php?key=%s&domain=%s", s.PassWord, s.Domain)
	s.http.New("GET", uri)
	s.http.Execute()
	defer s.http.Close()
	ret, _ := s.http.Byte()
	dmresult := DomainBoomResult{}
	if err := json.Unmarshal(ret, &dmresult); err != nil {
		log.Println("[!]GetResult Error: ", err)
	}
	for _, v := range dmresult {
		s.result[v] = "未解析"
	}
	return s.result
}

func (s *DomainBoom) Login(_ bool) bool {
	return true
}
