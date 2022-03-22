package Gsensor

import (
	"encoding/json"
	"fmt"
	"github.com/code-scan/Goal/Gconvert"
	"github.com/code-scan/Goal/Ghttp"
	"log"
)

type Beian struct {
	Domain string
	Type   string
	result Result
	http   Ghttp.Http
}

type BeianResult struct {
	Flag      bool        `json:"flag"`
	Message   string      `json:"message"`
	ErrorCode interface{} `json:"errorCode"`
	Result    ResultBeian `json:"result"`
}
type BeianContent struct {
	ID               string      `json:"id"`
	MainID           string      `json:"mainId"`
	UnitName         string      `json:"unitName"`
	MainLicence      string      `json:"mainLicence"`
	NatureID         string      `json:"natureId"`
	ServiceID        string      `json:"serviceId"`
	ServiceName      string      `json:"serviceName"`
	HomeURL          string      `json:"homeUrl"`
	LeaderName       interface{} `json:"leaderName"`
	Domain           string      `json:"domain"`
	UpdateRecordTime string      `json:"updateRecordTime"`
	Plist            interface{} `json:"plist"`
}
type ResultBeian struct {
	Content       []BeianContent `json:"content"`
	Sorted        interface{}    `json:"sorted"`
	TotalElements int            `json:"totalElements"`
	SizeOfContent int            `json:"sizeOfContent"`
	Paged         interface{}    `json:"paged"`
	First         interface{}    `json:"first"`
	Last          interface{}    `json:"last"`
	TotalPages    int            `json:"totalPages"`
	PageNumber    int            `json:"pageNumber"`
	PageSize      int            `json:"pageSize"`
}

func (s *Beian) GetInfo() string {
	return "Beian ver 0.1 with  " + s.Type

}

// 域名查询备案号和备案主题
// 备案主体查询 备案号和备案域名

func (s *Beian) SetDomain(domain string) {
	s.Domain = domain
}

func (s *Beian) SetAccount(_ string) {
	return
}

func (s *Beian) SetPassword(_ string) {
	return
}

func (s *Beian) SetType(type_ string) {
	s.Type = type_
}

func (s *Beian) GetResult() Result {
	s.result = Result{}

	switch s.Type {
	case "beian":
		s.Beian()
	}
	return s.result
}
func (s *Beian) Beian() {
	postData := fmt.Sprintf("keyword=%s&pageIndex=1&pageSize=2000", Gconvert.UrlEncode(s.Domain))
	s.http.Post("https://m-beian.miit.gov.cn/webrec/queryRec", postData)
	s.http.SetContentType("application/x-www-form-urlencoded")
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Byte()
	log.Println(string(ret))
	if err != nil {
		log.Println(err)
		return
	}
	var r BeianResult
	json.Unmarshal(ret, &r)
	for _, result := range r.Result.Content {
		key := fmt.Sprintf("%s|||%s", result.MainLicence, result.UnitName)
		if _, ok := s.result[key]; ok {
			s.result[key] = s.result[key] + ";" + result.Domain
		}
		s.result[key] = result.Domain
	}
}
func (s *Beian) Login(_ bool) bool {
	return true
}
