package Gsensor

import (
	"fmt"
	"log"

	"github.com/code-scan/Goal/Ghttp"
)

type BeianXCN struct {
	Domain string
	Type   string
	result Result
	http   Ghttp.Http
}

func (s *BeianXCN) GetInfo() string {
	return "BeianXCN ver 0.1 with  " + s.Type

}

// 域名查询备案号和备案主题
// 备案主体查询 备案号和备案域名

func (s *BeianXCN) SetDomain(domain string) {
	s.Domain = domain
}

func (s *BeianXCN) SetAccount(_ string) {
	return
}

func (s *BeianXCN) SetPassword(_ string) {
	return
}

func (s *BeianXCN) SetType(type_ string) {
	s.Type = type_
}

func (s *BeianXCN) GetResult() Result {
	s.result = Result{}

	switch s.Type {
	case "BeianXCN":
		s.BeianXCN()
	}
	return s.result
}
func (s *BeianXCN) BeianXCN() {
	uri := fmt.Sprintf("http://www.beianx.cn/search/%s", s.Domain)
	s.http.Get(uri)
	s.http.SetContentType("application/x-www-form-urlencoded")
	s.http.Execute()
	defer s.http.Close()
	ret, err := s.http.Byte()
	log.Println(string(ret))
	if err != nil {
		log.Println(err)
		return
	}
	// var r BeianXCNResult
	// json.Unmarshal(ret, &r)
	// for _, result := range r.Result.Content {
	// 	key := fmt.Sprintf("%s|||%s", result.MainLicence, result.UnitName)
	// 	if _, ok := s.result[key]; ok {
	// 		s.result[key] = s.result[key] + ";" + result.Domain
	// 		continue
	// 	}
	// 	s.result[key] = result.Domain
	// }
}
func (s *BeianXCN) Login(_ bool) bool {
	return true
}
