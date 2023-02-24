package Gsensor

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/code-scan/Goal/Ghttp"
)

type RapidDns struct {
	Domain string
	Type   string
	result Result
	http   Ghttp.Http
}

type RapidDnsResult struct {
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

func (s *RapidDns) GetInfo() string {
	return "RapidDnsImpl ver 0.1 with  " + s.Type

}

func (s *RapidDns) SetDomain(domain string) {
	s.Domain = domain
}

func (s *RapidDns) SetAccount(_ string) {
	return
}

func (s *RapidDns) SetPassword(_ string) {
	return
}

func (s *RapidDns) SetType(type_ string) {
	s.Type = type_
}

func (s *RapidDns) GetResult() Result {
	result := Result{}
	if s.Type != "subdomain" {
		return result
	}
	s.http.Get(fmt.Sprintf("https://rapiddns.io/subdomain/%s?full=1", s.Domain))
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Byte()
	if err != nil {
		log.Println("[!]GetResult Error: ", err, s.GetInfo())
		return result
	}
	var reader bytes.Buffer
	reader.Write(ret)
	doc, err := goquery.NewDocumentFromReader(&reader)
	if err != nil {
		return result
	}
	doc.Find("tr").Each(func(tri int, tr *goquery.Selection) {
		host := tr.Find("td:nth-child(2)").First().Text()
		ip := tr.Find("td:nth-child(3)").First().Text()
		host = strings.TrimSpace(host)
		ip = strings.TrimSpace(ip)
		result[host] = ip
	})

	s.result = result
	return result

}

func (s *RapidDns) Login(_ bool) bool {
	return true
}
