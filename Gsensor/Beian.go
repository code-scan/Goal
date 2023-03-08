package Gsensor

import (
	"encoding/json"
	"fmt"
	"log"

	"git.dev.me/jerry/Goal/Gconvert"
	"git.dev.me/jerry/Goal/Ghttp"
)

type Beian struct {
	Domain  string
	Type    string
	Account string
	result  Result
	http    Ghttp.Http
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

type BeianRet struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Params  Params `json:"params"`
	Success bool   `json:"success"`
}
type List struct {
	ContentTypeName  string `json:"contentTypeName"`
	Domain           string `json:"domain"`
	DomainID         int64  `json:"domainId"`
	LeaderName       string `json:"leaderName"`
	LimitAccess      string `json:"limitAccess"`
	MainID           int    `json:"mainId"`
	MainLicence      string `json:"mainLicence"`
	NatureName       string `json:"natureName"`
	ServiceID        int    `json:"serviceId"`
	ServiceLicence   string `json:"serviceLicence"`
	UnitName         string `json:"unitName"`
	UpdateRecordTime string `json:"updateRecordTime"`
}
type Params struct {
	EndRow           int    `json:"endRow"`
	FirstPage        int    `json:"firstPage"`
	HasNextPage      bool   `json:"hasNextPage"`
	HasPreviousPage  bool   `json:"hasPreviousPage"`
	IsFirstPage      bool   `json:"isFirstPage"`
	IsLastPage       bool   `json:"isLastPage"`
	LastPage         int    `json:"lastPage"`
	List             []List `json:"list"`
	NavigatePages    int    `json:"navigatePages"`
	NavigatepageNums []int  `json:"navigatepageNums"`
	NextPage         int    `json:"nextPage"`
	OrderBy          string `json:"orderBy"`
	PageNum          int    `json:"pageNum"`
	PageSize         int    `json:"pageSize"`
	Pages            int    `json:"pages"`
	PrePage          int    `json:"prePage"`
	Size             int    `json:"size"`
	StartRow         int    `json:"startRow"`
	Total            int    `json:"total"`
}

func (s *Beian) GetInfo() string {
	return "Beian ver 0.1 with  " + s.Type

}

// 域名查询备案号和备案主题
// 备案主体查询 备案号和备案域名

func (s *Beian) SetDomain(domain string) {
	s.Domain = domain
}

func (s *Beian) SetAccount(ac string) {
	s.Account = ac
}

func (s *Beian) SetPassword(_ string) {
}

func (s *Beian) SetType(type_ string) {
	s.Type = type_
}

func (s *Beian) GetResult() Result {
	s.result = Result{}
	if s.Account == "" {
		return s.result
	}
	switch s.Type {
	case "beian":
		s.Beian()
	}
	return s.result
}
func (s *Beian) Beian() {
	uri := fmt.Sprintf("%s/open?domain=%s", s.Account, Gconvert.UrlEncode(s.Domain))
	s.http.Get(uri)
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Byte()
	log.Println(string(ret))
	if err != nil {
		log.Println(err)
		return
	}
	var r BeianRet
	json.Unmarshal(ret, &r)
	log.Println(r)
	for _, result := range r.Params.List {
		key := fmt.Sprintf("%s|||%s", result.MainLicence, result.UnitName)
		if _, ok := s.result[key]; ok {
			s.result[key] = s.result[key] + ";" + result.Domain
			continue
		}
		s.result[key] = result.Domain
	}
}
func (s *Beian) Login(_ bool) bool {
	return true
}
