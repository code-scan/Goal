package Gsensor

import (
	"encoding/json"
	"fmt"
	"github.com/code-scan/Goal/Ghttp"
	"github.com/code-scan/Goal/Gnet"
	"log"
	"strconv"
	"strings"
)

type SecurityTrails struct {
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
}

var Cookie string

func (s *SecurityTrails) GetInfo() string {
	return "SecurityTrailsImpl ver 0.2 with " + s.Type
}
func (s *SecurityTrails) SetType(Type string) {
	s.Type = Type
}
func (s *SecurityTrails) SetDomain(domain string) {
	s.Domain = domain
}
func (s *SecurityTrails) SetAccount(account string) {
	s.UserName = account
}
func (s *SecurityTrails) SetPassword(password string) {
	s.PassWord = password
}
func (s *SecurityTrails) GetResult() Result {
	s.result = Result{}
	switch s.Type {
	case "subdomain":
		s.GetSubDomain(1)
	case "sameserver":
		s.GetIp(1)
	case "ahistory":
		s.GetHistory()
	}
	return s.result
}
func (s *SecurityTrails) CheckLogin() bool {
	s.http.New("GET", "https://securitytrails.com/app/account")
	s.http.Execute()
	ret, _ := s.http.Text()
	return strings.Contains(ret, s.UserName)
}
func (s *SecurityTrails) Login(ReLogin bool) bool {
	s.MaxPage = 100
	s.http = Ghttp.Http{}
	if ReLogin == false {
		s.http.SetCookie(s.Cookie)
		return true
	}
	if s.CheckLogin() {
		return true
	}
	postData := make(map[string]interface{})
	postData["email"] = s.UserName
	postData["password"] = s.PassWord

	s.http.New("POST", "https://securitytrails.com/api/auth/login")
	s.http.SetPostJson(postData)
	s.http.Execute()
	s.Cookie = s.http.RespCookie()
	s.GetBuildId()
	return true
}
func (s *SecurityTrails) GetBuildId() {
	ret, err := s.httpReq("https://securitytrails.com/list/apex_domain/baidu.com")
	if err != nil {
		log.Println("[!] ", s.GetInfo(), " GetBuildId Error: ", err)
	}
	resp := strings.Split(string(ret), `"buildId":"`)
	if len(resp) > 1 {
		s.buildId = strings.Split(resp[1], `"`)[0]
	}
}

//子域名查询
func (s *SecurityTrails) GetSubDomain(page int) {
	if page == 1 {
		s.result = Result{}
	}
	uri := "https://securitytrails.com/app/api/v1/list_new/hostname/" + s.Domain + "?page=" + strconv.Itoa(page)
	uri = fmt.Sprintf("https://securitytrails.com/_next/data/%s/list/apex_domain/%s.json?domain=%s&page=%d", s.buildId, s.Domain, s.Domain, page)
	resp, err := s.httpReq(uri)

	if err != nil {
		log.Println("[!] ", s.GetInfo(), " GetSubDomain 1 Error: ", err)
		return
	}
	hostName := SubDomainRespsonse{}
	err = json.Unmarshal(resp, &hostName)
	if err != nil {
		log.Println("[!] ", s.GetInfo(), " GetSubDomain 2 Error: ", err)
		return

	}
	for _, r := range hostName.Pageprops.Apexdomaindata.Data.Records {
		if s.Resolver != true {
			s.result[r.Hostname] = "未解析"
			continue
		}
		ips := Gnet.GetHostIp(r.Hostname, "8.8.8.8:53")
		if len(ips) > 0 {
			s.result[r.Hostname] = ips[0]
		} else {
			s.result[r.Hostname] = "未解析"
		}
	}
	if page >= s.MaxPage {
		return
	}
	if hostName.Pageprops.Apexdomaindata.Data.Total != len(s.result) {
		s.GetSubDomain(page + 1)
	}
}

//历史记录查询
func (s *SecurityTrails) GetHistory() {
	s.result = Result{}
	uri := "https://securitytrails.com/app/api/v1/history/" + s.Domain + "/dns/a"
	uri = fmt.Sprintf("https://securitytrails.com/_next/data/%s/domain/%s/history/a.json?domain=%s&type=a", s.buildId, s.Domain, s.Domain)
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
func (s *SecurityTrails) GetIp(page int) {
	if page == 1 {
		s.result = Result{}
	}
	uri := "https://securitytrails.com/app/api/v1/list_new/ip/" + s.Domain + "?page=" + strconv.Itoa(page)
	uri = fmt.Sprintf("https://securitytrails.com/_next/data/%s/list/ip/%s.json?page=%d&ip=%s", s.buildId, s.Domain, page, s.Domain)
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
func (s *SecurityTrails) httpReq(uri string) ([]byte, error) {
	s.http.New("POST", uri)
	s.http.SetCookie(s.Cookie)
	s.http.Execute()
	return s.http.Byte()
}
