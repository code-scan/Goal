package Gsensor

import (
	"encoding/json"
	"fmt"
	"github.com/code-scan/Goal/Gconvert"
	"github.com/code-scan/Goal/Ghttp"
	"log"
)

type Shodan struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	http     Ghttp.Http
}
type ShodanResult struct {
	RegionCode   interface{}   `json:"region_code"`
	Tags         []interface{} `json:"tags"`
	IP           int           `json:"ip"`
	AreaCode     interface{}   `json:"area_code"`
	Domains      []interface{} `json:"domains"`
	Hostnames    []interface{} `json:"hostnames"`
	PostalCode   interface{}   `json:"postal_code"`
	DmaCode      interface{}   `json:"dma_code"`
	CountryCode  string        `json:"country_code"`
	Org          string        `json:"org"`
	Data         []interface{} `json:"data"`
	Asn          string        `json:"asn"`
	City         interface{}   `json:"city"`
	Latitude     float64       `json:"latitude"`
	Isp          string        `json:"isp"`
	Longitude    float64       `json:"longitude"`
	LastUpdate   string        `json:"last_update"`
	CountryCode3 interface{}   `json:"country_code3"`
	CountryName  string        `json:"country_name"`
	IPStr        string        `json:"ip_str"`
	Os           interface{}   `json:"os"`
	Ports        []int         `json:"ports"`
}

func (s *Shodan) GetInfo() string {
	return "ShodanImpl ver 0.1 with " + s.Type
}
func (s *Shodan) SetType(Type string) {
	s.Type = Type
}

func (s *Shodan) SetDomain(domain string) {
	s.Domain = domain
}
func (s *Shodan) SetAccount(account string) {
	s.UserName = account
}
func (s *Shodan) SetPassword(password string) {
	s.PassWord = password
}
func (s *Shodan) GetResult() Result {
	var result = Result{}
	if s.Type != "ports" {
		return result
	}
	uri := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=MM72AkzHXdHpC8iP65VVEEVrJjp7zkgd&minify=true", s.Domain)
	s.http.New("GET", uri)
	s.http.Execute()
	ret, _ := s.http.Byte()
	shodanResult := ShodanResult{}
	err := json.Unmarshal(ret, &shodanResult)
	if err != nil {
		log.Println("Shodan GetResult Error: ", err, " Domain: ", s.Domain)
		return result
	}
	for _, v := range shodanResult.Ports {
		port := Gconvert.Int2String(v)
		result[s.Domain+":"+port] = s.Domain
	}
	return result
}
func (s *Shodan) Login(ReLogin bool) bool {

	return true
}
