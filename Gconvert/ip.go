package Gconvert

import (
	"encoding/json"
	"fmt"
	"log"

	"git.dev.me/jerry/Goal/Ghttp"
)

type IPinfoResult struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func IPinfo(ip string, token string) IPinfoResult {
	var result IPinfoResult
	http := Ghttp.Http{}
	http.New("GET", fmt.Sprintf("https://ipinfo.io/%s/json", ip))
	if token != "" {
		http.SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	http.Execute()
	defer http.Close()
	ret, err := http.Byte()
	if err != nil {
		log.Println("[!] IPinfo Error: ", err)
		return result
	}
	if err = json.Unmarshal(ret, &result); err != nil {
		log.Println("[!] IPinfo Unmarshal Error: ", err)
	}
	return result

}
