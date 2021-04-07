package Gconvert

import (
	"log"
	"strconv"
)

// int float转string
func Int2String(i interface{}) string {
	switch i.(type) {
	case int:
		return strconv.FormatInt(int64(i.(int)), 10)
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	}
	return ""

}

func Str2Int(s string) int {
	ret, err := strconv.Atoi(s)
	if err != nil {
		log.Println("[!] Str2Int Error: ", err)

		return 0
	}
	return ret
}
func Str2Int64(s string) int64 {
	ret, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("[!] Str2Int64 Error: ", err)

		return 0
	}
	return ret
}

func Str2Float(s string) float32 {
	ret, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Println("[!] Str2Float Error: ", err)

		return 0
	}
	return float32(ret)
}
func Str2Float64(s string) float64 {
	ret, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("[!] Str2Float64 Error: ", err)

		return 0
	}
	return ret
}
