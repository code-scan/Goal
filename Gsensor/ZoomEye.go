package Gsensor

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"git.dev.me/jerry/Goal/Ghttp"
)

type ZoomEye struct {
	Domain   string
	UserName string
	PassWord string
	Type     string
	result   Result
	http     Ghttp.Http
}

type ZoomEyeResult struct {
	Status int                 `json:"status"`
	Total  int                 `json:"total"`
	List   []ZoomEyeResultList `json:"list"`
	Msg    string              `json:"msg"`
	Type   int                 `json:"type"`
}
type ZoomEyeResultList struct {
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func (s *ZoomEye) GetInfo() string {
	return "ZoomEye ver 0.1 with  " + s.Type
}
func (s *ZoomEye) SetType(Type string) {
	s.Type = Type
}
func (s *ZoomEye) SetDomain(domain string) {
	s.Domain = domain
}
func (s *ZoomEye) SetAccount(account string) {
	s.UserName = account
}
func (s *ZoomEye) SetPassword(password string) {
	s.PassWord = password
}
func (s *ZoomEye) GetResult() Result {
	s.result = Result{}
	switch s.Type {
	case "subdomain":
		s.GetSubDomain(1)
	case "sameserver":
		s.GetIp(1)
	case "ports":
		break
	}
	return s.result
}
func (s *ZoomEye) Login(ReLogin bool) bool {
	//s.http = HttpHelper{}
	return true
}

func (s *ZoomEye) GetSubDomain(i int) {
	//s.result = Result{}
	resp := s.send(fmt.Sprintf(`q=%s&type=1&page=%d&s=500`, s.Domain, i))
	zoomeyeresult := ZoomEyeResult{}
	json.Unmarshal(resp, &zoomeyeresult)
	for _, v := range zoomeyeresult.List {
		host := v.Name
		if strings.Contains(host, "://") {
			host = strings.Split(host, "://")[1]
		}
		s.result[host] = v.IP
	}
	if zoomeyeresult.Status != 200 {
		return
	}
	if len(s.result) < zoomeyeresult.Total {
		s.GetSubDomain(i + 1)
	}

}

func (s *ZoomEye) GetIp(i int) {
	//s.result = Result{}
	// q=172.67.168.89&type=1&page=0
	resp := s.send(fmt.Sprintf(`q=%s&type=0&page=%d&size=500`, s.Domain, i))
	zoomeyeresult := ZoomEyeResult{}
	json.Unmarshal(resp, &zoomeyeresult)
	for _, v := range zoomeyeresult.List {
		host := v.Name
		if strings.Contains(host, "://") {
			host = strings.Split(host, "://")[1]
		}
		s.result[host] = v.IP
	}
	if zoomeyeresult.Status != 200 {
		return
	}
	if len(s.result) < zoomeyeresult.Total {
		s.GetIp(i + 1)
	}
}

func (s *ZoomEye) GetPorts() {

}

func (s *ZoomEye) send(query string) []byte {
	var result []byte
	//query = base64.StdEncoding.EncodeToString([]byte(query))
	urls := fmt.Sprintf("https://api.zoomeye.org/domain/search?%s", query)
	s.http.New("GET", urls)
	s.http.SetHeader("API-KEY", s.PassWord)
	s.http.Execute()
	defer s.http.Close()
	resp, err := s.http.Byte()
	if err != nil {
		log.Println(err)
		return result
	}
	return resp
}
