package Gsensor

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/code-scan/Goal/Ghttp"
	"github.com/code-scan/Goal/Gnet"
)

type SecurityTrailsApi struct {
	Domain   string
	UserName string
	PassWord string
	Cookie   string
	Type     string
	Resolver bool
	MaxPage  int
	http     Ghttp.Http
	result   Result
	buildId  string
	hostIp   string
}

type SecurityTrailsApiResponse struct {
	Endpoint string `json:"endpoint"`
	Meta     struct {
		LimitReached bool `json:"limit_reached"`
	} `json:"meta"`
	SubdomainCount int      `json:"subdomain_count"`
	Subdomains     []string `json:"subdomains"`
}

func (s *SecurityTrailsApi) GetInfo() string {
	return "SecurityTrailsApiImpl ver 0.2 with " + s.Type
}
func (s *SecurityTrailsApi) SetType(Type string) {
	s.Type = Type
}
func (s *SecurityTrailsApi) SetDomain(domain string) {
	s.Domain = domain
}
func (s *SecurityTrailsApi) SetAccount(account string) {
	s.UserName = account
}
func (s *SecurityTrailsApi) SetPassword(password string) {
	s.PassWord = password
}
func (s *SecurityTrailsApi) GetResult() Result {
	s.result = Result{}
	switch s.Type {
	case "subdomain":
		s.GetSubDomain(1)

		// case "sameserver":
		// 	s.GetIp(1)
		// case "ahistory":
		// 	s.GetHistory()
	}
	return s.result
}

func (s *SecurityTrailsApi) GetBuildId() {
	ret, err := s.httpReq("https://SecurityTrailsApi.com/list/apex_domain/baidu.com")
	s.http.IgnoreSSL()
	s.http.SetHeader("Host", "SecurityTrailsApi.com")
	if err != nil {
		log.Println("[!] ", s.GetInfo(), " GetBuildId Error: ", err)
	}
	resp := strings.Split(string(ret), `"buildId":"`)
	if len(resp) > 1 {
		s.buildId = strings.Split(resp[1], `"`)[0]
	}
}

// 子域名查询
func (s *SecurityTrailsApi) GetSubDomain(page int) {
	if page == 1 {
		s.result = Result{}
	}
	uri := "https://api.securitytrails.com/v1/domain/%s/subdomains"
	uri = fmt.Sprintf(uri, s.Domain)
	resp, err := s.httpReq(uri)
	log.Println(string(resp))
	if err != nil {
		log.Println("[!] ", s.GetInfo(), " GetSubDomain 1 Error: ", err)
		return
	}
	hostName := SecurityTrailsApiResponse{}
	err = json.Unmarshal(resp, &hostName)
	if err != nil {
		log.Println(string(resp))
		log.Println("[!] ", s.GetInfo(), " GetSubDomain 2 Error: ", err)
		return

	}
	for _, r := range hostName.Subdomains {
		hostname := fmt.Sprintf("%s.%s", r, s.Domain)
		if !s.Resolver {
			s.result[hostname] = "未解析"
			continue
		}
		ips := Gnet.GetHostIp(hostname, "8.8.8.8:53")
		if len(ips) > 0 {
			s.result[hostname] = ips[0]
		} else {
			s.result[hostname] = "未解析"
		}
	}
}

// 历史记录查询
func (s *SecurityTrailsApi) GetHistory() {
	s.result = Result{}
	uri := "https://SecurityTrailsApi.com/app/api/v1/history/" + s.Domain + "/dns/a"
	uri = fmt.Sprintf("https://SecurityTrailsApi.com/_next/data/%s/domain/%s/history/a.json?domain=%s&type=a", s.buildId, s.Domain, s.Domain)
	resp, err := s.httpReq(uri)
	if err != nil {
		log.Println("[!] GetHistory 1 Error: ", err)
		return
	}
	history := AHistoryRespsonse{}
	err = json.Unmarshal(resp, &history)
	if err != nil {
		log.Println("[!] GetHistory 2 Error: ", err)
		return
	}
	for _, r := range history.Pageprops.Dnsdata.Data.History.A.Data.Records {
		for _, v := range r.Values {
			s.result[v.IP] = "[" + r.Firstseen + "] -- [" + r.Lastseen + "]"
		}
	}
}

// ip查询
func (s *SecurityTrailsApi) GetIp(page int) {
	if page == 1 {
		s.result = Result{}
	}
	uri := "https://SecurityTrailsApi.com/app/api/v1/list_new/ip/" + s.Domain + "?page=" + strconv.Itoa(page)
	uri = fmt.Sprintf("https://SecurityTrailsApi.com/_next/data/%s/list/ip/%s.json?page=%d&ip=%s", s.buildId, s.Domain, page, s.Domain)
	resp, err := s.httpReq(uri)
	if err != nil {
		log.Println("[!] GetIp 1 Error: ", err)
		return

	}
	ip := SameServerResponse{}
	err = json.Unmarshal(resp, &ip)
	if err != nil {
		log.Println("[!] GetIp 2 Error: ", err)
		return
	}
	for _, r := range ip.Pageprops.Serverresponse.Data.Records {
		s.result[r.Hostname] = s.Domain
	}
	if page >= s.MaxPage {
		return
	}
	if ip.Pageprops.Serverresponse.Data.Total != len(s.result) {
		s.GetIp(page + 1)
	}
}
func (s *SecurityTrailsApi) httpReq(uri string) ([]byte, error) {
	s.http.New("GET", uri)
	s.http.IgnoreSSL()
	s.http.SetCookie(s.Cookie)
	s.http.SetHeader("apikey", s.PassWord)
	s.http.Execute()
	log.Println(s.PassWord)
	defer s.http.Close()
	return s.http.Byte()
}
