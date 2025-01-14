package test

import (
	"log"
	"testing"

	"github.com/code-scan/Goal/Gsensor"
)

func TestFofa(t *testing.T) {
	ff := Gsensor.Fofa{}
	// api的邮箱与api的key
	ff.SetAccount("@qq.com")
	ff.SetPassword("s")
	ff.SetType("sameserver")
	ff.SetType("subdomain")
	//ff.SetType("ports")

	ff.SetDomain("freebuf.com")
	r := ff.GetResult()
	log.Println(r)
}

func TestBeian(t *testing.T) {
	beian := Gsensor.Beian{}
	beian.SetDomain("baidu.com")
	beian.SetAccount("http://127.0.0.1:65511/")
	beian.SetType("beian")
	result := beian.GetResult()
	log.Printf("%#v", result)
}

func TestSecTrail(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ss := Gsensor.SecurityTrails{}
	//可以不登录，只能查询第一页
	ss.SetAccount("xxxxx@gmail.com")
	ss.SetPassword("xxxxxxx")
	ss.Login(true)
	ss.MaxPage = 10 // 自定义最大翻页，登录后默认100页
	ss.SetDomain("360.cn")
	ss.SetType("subdomain")
	//ss.SetType("ahistory")

	//ss.SetDomain("172.67.168.89")
	//ss.SetType("sameserver")
	r := ss.GetResult()
	log.Println(len(r))
	log.Println(r)

}
func TestSecTrailApi(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ss := Gsensor.SecurityTrailsApi{}
	//可以不登录，只能查询第一页
	ss.SetPassword("NQIcGiQBA53myDkCS8wXj2d4MzauIdkH")
	ss.SetDomain("360.cn")
	ss.SetType("subdomain")
	r := ss.GetResult()
	log.Println(len(r))
	log.Println(r)

}
func TestZoom(t *testing.T) {
	z := Gsensor.ZoomEye{}
	z.SetType("subdomain")
	z.SetDomain("weibo.com")
	z.SetPassword("xxx-xxx-xxx-xxx-xxx")
	r := z.GetResult()
	log.Println(r)
}
func TestShodan(t *testing.T) {
	sd := Gsensor.Shodan{}
	// shodan用的chrome插件的api，只能用来获取端口
	sd.SetDomain("172.67.168.89")
	sd.SetType("ports")
	rr := sd.GetResult()
	log.Println(rr)
}
func TestBuff(t *testing.T) {
	bf := Gsensor.Bufferover{}
	bf.SetType("subdomain")
	bf.SetDomain("baidu.com")
	r := bf.GetResult()
	log.Println(r)
}

func TestRapid(t *testing.T) {
	bf := Gsensor.RapidDns{}
	bf.SetType("subdomain")
	bf.SetDomain("freebuf.com")
	r := bf.GetResult()
	log.Println(r)
}
func TestCrt(t *testing.T) {
	bf := Gsensor.CrtSh{}
	bf.SetType("subdomain")
	bf.SetDomain("baidu.com")
	r := bf.GetResult()
	log.Println(r)
}

func TestDomainBoom(t *testing.T) {
	bf := Gsensor.DomainBoom{}
	for {
		bf.SetPassword("xxxxxx")
		bf.SetDomain("freebuf.com")
		bf.SetType("subdomain")
		r := bf.GetResult()
		log.Println(r)
	}
}

func TestWebapp(t *testing.T) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	//bf := Gsensor.WappalyzerGo{}
	//bf.SetType("finger")
	//bf.SetDomain("http://baidu.com")
	//r := bf.GetResult()
	//log.Println(r)
	//return
	//var h http.Transport
	//h = http.Transport{}
	bf := Gsensor.WappalyzerGo{}
	for i := 0; i < 10000; i++ {
		bf.SetType("finger")
		bf.SetDomain("https://www.freebuf.com")
		//bf.Http.HttpTransport = &h
		r := bf.GetResult()
		log.Println(r)
	}
	select {}

}
func TestMas(t *testing.T) {
	mas := Gsensor.MassScan{}
	mas.SetType("ports")
	mas.SetDomain("1.1.1.1")
	mas.GetResult()
}
func TestQiYe(t *testing.T) {
	qiye := Gsensor.AiQiCha{}
	qiye.SetType("qiye")
	qiye.SetDomain("中国电信集团有限公司")
	qiye.SetAccount("")
	ret := qiye.GetResult()
	log.Println(ret)
}

func TestQiYeHold(t *testing.T) {
	qiye := Gsensor.AiQiCha{}
	qiye.SetType("qiye_hold")
	qiye.SetDomain("28696963178919")
	qiye.SetAccount("AiQiCha cookie")
	ret := qiye.GetResult()
	log.Println(ret)
}
func TestQuake(t *testing.T) {
	quake := Gsensor.Quake{}
	quake.SetPassword("40fc40e2-ff0a-4487-8e7e-a833ba9291b1")
	quake.SetAccount("as")
	quake.SetType("subdomain")
	quake.SetDomain("freebuf.com")
	// quake.GetSubDomain(1)
	log.Println(quake.GetResult())
}

// func TestSub(t *testing.T) {
// 	log.SetFlags(log.Lshortfile | log.LstdFlags)
// 	sub := Gsensor.NewKSubDomainSensor()
// 	sub.SetType("freebuf.com")
// 	sub.SetType("subdomain")
// 	log.Println(sub.GetResult())
// }
