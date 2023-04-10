package test

import (
	"log"
	"testing"
	"time"

	"github.com/code-scan/Goal/Gconvert"
)

func TestConvert(t *testing.T) {

	var i = "1234567"
	var f = "123456.2345"
	var ff = 12345.6789
	var ii = 1234567

	log.Println(ii, Gconvert.Int2String(ii))
	log.Println(ff, Gconvert.Int2String(ff))

	log.Println(i, Gconvert.Str2Int(i))
	log.Println(f, Gconvert.Str2Int(f))

	log.Println(f, Gconvert.Str2Float(f))
	log.Println(f, Gconvert.Str2Float64(f))

	log.Println(i, Gconvert.Str2Float(i))
	log.Println(i, Gconvert.Str2Float64(i))

	log.Println(Gconvert.Str2Url("1"))
	log.Println("encode base64 ", Gconvert.B64Encode("12312312"))
	log.Println("decode base64 ", Gconvert.B64Decode("324 d"))

	log.Println("urlencode ", Gconvert.UrlEncode("324=1;sd;'123 d"))
	log.Println("urldecode  ", Gconvert.UrlDecode("%25%27%22"))
	log.Println("rawurl ", Gconvert.RawDecode("%25%27%22"))
	log.Println("raw encode  ", Gconvert.RawEncode("324=1;sd;'123 d"))

	log.Println("Time2Str ", Gconvert.Time2Str(time.Now()))
	log.Println("Unix2Time ", Gconvert.Unix2Time(1614168000))
	log.Println("Str2Time ", Gconvert.Str2Time("2020-112-11 22:33:11"))

	log.Println("md5 ", Gconvert.Md5("123456"))
	log.Println("sha1 ", Gconvert.Sha1("123456"))
	log.Println("sha256 ", Gconvert.Sha256("123456"))
	log.Println("sha512 ", Gconvert.Sha512("123456"))

	log.Println("ipinfo ", Gconvert.IPinfo("104.21.78.188", "ipinfo key"))

}
