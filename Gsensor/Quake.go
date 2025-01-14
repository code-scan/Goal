package Gsensor

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/code-scan/Goal/Ghttp"
)

type Quake struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	result   Result
	http     Ghttp.Http
}

type QuakeResult struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    []UtilsFunData `json:"data"`
	// Meta    UtilsFunMeta   `json:"meta"`
}

type UtilsFunData struct {
	// Components []UtilsFunDataComponents `json:"components"`
	// Images     []UtilsFunDataImages     `json:"images"`
	Org       string `json:"org"`
	Ip        string `json:"ip"`
	Is_ipv6   bool   `json:"is_ipv6"`
	Transport string `json:"transport"`
	Hostname  string `json:"hostname"`
	Port      int    `json:"port"`
	// Service    UtilsFunDataService      `json:"service"`
	Domain string `json:"domain"`
	// Location   UtilsFunDataLocation     `json:"location"`
	Time string `json:"time"`
	Asn  int    `json:"asn"`
	Id   string `json:"id"`
}

func (s *Quake) GetInfo() string {
	return "QuakeImpl ver 0.1 with  " + s.Type
}
func (s *Quake) SetType(Type string) {
	s.Type = Type
}
func (s *Quake) SetDomain(domain string) {
	s.Domain = domain
}
func (s *Quake) SetAccount(account string) {
	s.UserName = account
}
func (s *Quake) SetPassword(password string) {
	s.PassWord = password
}
func (s *Quake) GetResult() Result {
	s.result = Result{}
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
func (s *Quake) Login(ReLogin bool) bool {
	//s.http = HttpHelper{}
	return true
}

func (s *Quake) GetSubDomain(i int) {
	s.result = Result{}
	resp := s.send(`domain:"` + s.Domain + `"`)
	// log.Println(string(resp))
	Quakeresult := QuakeResult{}
	json.Unmarshal(resp, &Quakeresult)
	for _, v := range Quakeresult.Data {
		host := v.Domain
		if strings.Contains(host, "://") {
			host = strings.Split(host, "://")[1]
		}
		s.result[host] = v.Ip
	}

}

func (s *Quake) GetIp(i int) {
	s.result = Result{}
	// resp := s.send(`ip="` + s.Domain + `" && type=subdomain`)
	// Quakeresult := QuakeResult{}
	// json.Unmarshal(resp, &Quakeresult)
	// for _, v := range Quakeresult.Results {
	// 	host := v[0]
	// 	if strings.Contains(host, "://") {
	// 		host = strings.Split(host, "://")[1]
	// 	}
	// 	s.result[host] = v[1]
	// }
}

func (s *Quake) GetPorts() {
	s.result = Result{}
	// resp := s.send(`ip="` + s.Domain + `" && type=service`)
	// Quakeresult := QuakeResult{}
	// json.Unmarshal(resp, &Quakeresult)
	// for _, v := range Quakeresult.Results {
	// 	host := v[0]
	// 	if strings.Contains(host, "://") {
	// 		host = strings.Split(host, "://")[1]
	// 	}
	// 	s.result[host] = v[1]
	// }
}

// 40fc40e2-ff0a-4487-8e7e-a833ba9291b1
func (s *Quake) send(query string) []byte {
	var result []byte
	// query = base64.StdEncoding.EncodeToString([]byte(query))
	urls := "https://quake.360.net/api/v3/search/quake_service"
	// log.Println(urls)
	s.http.New("POST", urls)
	s.http.SetHeader("X-QuakeToken", s.PassWord)
	s.http.SetHeader("Content-Type", "application/json")
	query = strings.ReplaceAll(query, "\"", "\\\"")
	query = `{
		"query": "` + query + `",
		"start": 0,
		"size": 10000
	}`
	// log.Println(query)
	// return nil
	s.http.SetPostString(query)
	s.http.Execute()
	defer s.http.Close()
	resp, err := s.http.Byte()
	if err != nil {
		log.Println(err)
		return result
	}
	return resp
}
