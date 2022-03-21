package Gsensor

import (
	"encoding/json"
	"fmt"
	"github.com/CUCyber/ja3transport"
	"github.com/code-scan/Goal/Gconvert"
	"github.com/code-scan/Goal/Ghttp"
	tls "github.com/refraction-networking/utls"
	"log"
	"strings"
)

type AiQiCha struct {
	Domain string
	Type   string
	result Result
	Cookie string
	http   Ghttp.Http
}
type Browser struct {
	Ja3       string
	UserAgent string
}

func (s *AiQiCha) GetInfo() string {
	return "AiQiCha ver 0.1 with  " + s.Type

}

// 这里传入公司名字
func (s *AiQiCha) SetDomain(domain string) {
	s.Domain = domain
}

func (s *AiQiCha) SetAccount(cookie string) {
	s.Cookie = cookie
	return
}

func (s *AiQiCha) SetPassword(_ string) {
	return
}

func (s *AiQiCha) SetType(type_ string) {
	s.Type = type_
}

func (s *AiQiCha) GetResult() Result {
	s.result = Result{}

	switch s.Type {
	case "qiye":
		s.getQiye()
	case "qiye_hold":
		s.getQiyeHold()
	}

	// 获取子公司

	return s.result

}

func (s *AiQiCha) getQiye() {
	result := s.searhKeyword()
	log.Println(result)
	for _, corp := range result.Data.ResultList {
		corp.EntName = strings.ReplaceAll(corp.EntName, "<em>", "")
		corp.EntName = strings.ReplaceAll(corp.EntName, "</em>", "")
		corp.EntName = fmt.Sprintf("%s|||%s", corp.EntName, corp.Pid)
		s.result[corp.EntName] = s.getDomain(corp.Pid)
	}
}
func (s *AiQiCha) getQiyeHold() Result {
	// 通过关键词查询，这里讲道理应该传递进来的企业id会更简单
	uri := fmt.Sprintf("https://aiqicha.baidu.com/detail/holdsAjax?pid=%s&page=1&size=100", s.Domain)
	ret, _ := s.get(uri)
	var resp AiQiChaHoldResponse
	json.Unmarshal(ret, &resp)
	for _, l := range resp.Data.List {
		s.result[l.EntName] = s.getDomain(l.Pid)
	}
	return s.result

}
func (s *AiQiCha) searhKeyword() AiQiChaSearchResponse {
	uri := fmt.Sprintf("https://aiqicha.baidu.com/app/advanceFilterAjax?o=0&p=1&q=%s&t=111", Gconvert.UrlEncode(s.Domain))
	ret, _ := s.get(uri)
	var response AiQiChaSearchResponse
	json.Unmarshal(ret, &response)
	return response
}
func (s *AiQiCha) getDomain(id string) string {
	uri := fmt.Sprintf("https://aiqicha.baidu.com/appcompdata/headinfoAjax?pid=%s", id)
	ret, _ := s.get(uri)
	var response AiQiChaInfoResponse
	json.Unmarshal(ret, &response)
	website := response.Data.Website
	icp := s.getICP(id)
	if icp != "" {
		if website != "" {
			website = website + ";" + icp
		} else {
			return icp
		}
	}
	return website
}

func (s *AiQiCha) getICP(id string) string {
	uri := fmt.Sprintf("https://aiqicha.baidu.com/m/icpInfoAjax?pid=%s&page=0&size=100", id)
	ret, _ := s.get(uri)
	var resp AiQiChaICPResponse
	json.Unmarshal(ret, &resp)
	website := ""
	for _, l := range resp.Data.List {
		if len(l.Domain) > 0 {
			if website != "" {
				website = website + ";"
			}
			website = website + strings.Join(l.Domain, ";")
		}
	}
	return website
}

func (s *AiQiCha) get(uri string) ([]byte, error) {
	s.http.New("GET", uri)
	s.http.SetHeader("Zx-Open-Url", "https://aiqicha.baidu.com")
	s.http.SetHeader("Cuid", "1BA3566BA8CA3AFC39D4B7689710F79680C5707DEFBNFGSIODD")
	s.http.SetHeader("X-Requested-With", "XMLHttpRequest")
	s.http.SetHeader("Env", "IOS")
	s.http.SetHeader("Client-Version", "2.4.7")
	s.http.SetHeader("Host", "aiqicha.baidu.com")
	s.http.SetHeader("User-Agent", "aiinquiry/2.4.7 (iPhone; iOS 15.4; Scale/3.00) Ios (21) 11462417584689898958 aiqicha/2.4.7")
	s.http.SetHeader("Referer", "https://aiqicha.baidu.com/usercenter")
	s.http.SetHeader("Connection", "keep-alive")
	//s.http.IgnoreSSL()
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	ja3List := []Browser{
		{Ja3: "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53-10,0-23-65281-10-11-35-16-5-13-18-51-45-43-27,29-23-24,0", UserAgent: "aiinquiry/2.4.7 (iPhone; iOS 15.4; Scale/3.00) Ios (21) 11462417584689898958 aiqicha/2.4.7"},
		{Ja3: "771,4865-4866-4867-49196-49195-52393-49200-49199-52392-49188-49187-49162-49161-49192-49191-49172-49171-157-156-61-60-53-47-49160-49170-10,0-23-65281-10-11-16-5-13-18-51-45-43-21,29-23-24-25,0", UserAgent: "aiinquiry/2.4.7 (iPhone; iOS 15.4; Scale/3.00) Ios (21) 11462417584689898958 aiqicha/2.4.7"},
		{Ja3: "771,4865-4866-4867-49196-49195-49188-49187-49162-49161-52393-49200-49199-49192-49191-49172-49171-52392-157-156-61-60-53-47-49160-49170-10,65281-0-23-13-5-18-16-11-51-45-43-10-21,29-23-24-25,0", UserAgent: "aiinquiry/2.4.7 (iPhone; iOS 15.4; Scale/3.00) Ios (21) 11462417584689898958 aiqicha/2.4.7"},
	}
	tr, _ := ja3transport.NewTransportWithConfig(ja3List[2].Ja3, &config)
	s.http.HttpTransport = tr
	s.http.SetCookie(s.Cookie)
	//s.http.SetProxy("http://127.0.0.1:8080")
	s.http.HttpClient.Transport = tr
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Byte()
	return ret, err
}

func (s *AiQiCha) Login(_ bool) bool {
	return true
}
