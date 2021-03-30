package Gconvert

import (
	"encoding/base64"
	"log"
)

func B64Encode(param interface{}) string {
	var ret string
	switch param.(type) {
	case []byte:
		ret = base64.StdEncoding.EncodeToString(param.([]byte))
	case string:
		ret = base64.StdEncoding.EncodeToString([]byte(param.(string)))
	}
	return ret
}
func B64Decode(param string) string {
	ret, err := base64.StdEncoding.DecodeString(param)
	if err != nil {
		log.Println("[!] b64decode Error: ", err)
		return ""
	}
	return string(ret)
}
