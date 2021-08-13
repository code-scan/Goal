package Gsensor

import (
	"encoding/json"
	"github.com/code-scan/Goal/Ghttp"
	"log"
	"strings"
)

type Bufferover struct {
	Domain string
	Type   string
	result Result
	http   Ghttp.Http
}

type BufferoverResult struct {
	Meta struct {
		Runtime   string   `json:"Runtime"`
		Errors    []string `json:"Errors"`
		Message   string   `json:"Message"`
		Filenames []string `json:"FileNames"`
		Tos       string   `json:"TOS"`
	} `json:"Meta"`
	FdnsA []string    `json:"FDNS_A"`
	Rdns  interface{} `json:"RDNS"`
}

func (s *Bufferover) GetInfo() string {
	return "BufferoverImpl ver 0.1 with  " + s.Type

}

func (s *Bufferover) SetDomain(domain string) {
	s.Domain = domain
}

func (s *Bufferover) SetAccount(_ string) {
	return
}

func (s *Bufferover) SetPassword(_ string) {
	return
}

func (s *Bufferover) SetType(type_ string) {
	s.Type = type_
}

func (s *Bufferover) GetResult() Result {
	result := Result{}
	if s.Type != "subdomain" {
		return result
	}
	s.http.Get("https://dns.bufferover.run/dns?q=." + s.Domain)
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Byte()
	if err != nil {
		log.Println("[!]GetResult Error: ", err, s.GetInfo())
		return result
	}
	var buff BufferoverResult
	err = json.Unmarshal(ret, &buff)
	if err != nil {
		log.Println("[!]GetResult 2 Error: ", err, s.GetInfo())
		return result
	}
	for _, v := range buff.FdnsA {
		kv := strings.Split(v, ",")
		if len(kv) < 2 {
			continue
		}
		result[kv[1]] = kv[0]
	}
	s.result = result
	return result

}

func (s *Bufferover) Login(_ bool) bool {
	return true
}
