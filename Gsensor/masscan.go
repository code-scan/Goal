package Gsensor

import (
	"fmt"
	"github.com/code-scan/Goal/Ghttp"
	"github.com/dean2021/go-masscan"
	"log"
)

type MassScan struct {
	Domain string
	Type   string
	result Result
	http   Ghttp.Http
}

func (s *MassScan) GetInfo() string {
	return "MassScan ver 0.1 with  " + s.Type

}

func (s *MassScan) SetDomain(domain string) {
	s.Domain = domain
}

func (s *MassScan) SetAccount(_ string) {
	return
}

func (s *MassScan) SetPassword(_ string) {
	return
}

func (s *MassScan) SetType(type_ string) {
	s.Type = type_
}

func (s *MassScan) GetResult() Result {
	s.result = Result{}
	if s.Type != "ports" {
		return s.result
	}

	m := masscan.New()

	// masscan可执行文件路径,默认不需要设置
	//m.SetSystemPath("/usr/local/masscan/bin/masscan")

	// 扫描端口范围
	m.SetPorts("0-50")

	// 扫描IP范围
	m.SetRanges(s.Domain)

	// 扫描速率
	m.SetRate("2000")

	// 隔离扫描名单
	m.SetExclude("127.0.0.1")

	// 开始扫描
	err := m.Run()
	if err != nil {
		fmt.Println("scanner failed:", err)
		return s.result
	}

	// 解析扫描结果
	results, err := m.Parse()
	if err != nil {
		fmt.Println("Parse scanner result:", err)
		return s.result
	}

	for _, result := range results {
		log.Printf("%#v", result)
		for _, port := range result.Ports {
			key := fmt.Sprintf("%s:%s", result.Address.Addr, port.Portid)
			s.result[key] = port.State.State
		}
	}
	log.Println(s.result)
	return s.result
}

func (s *MassScan) Login(_ bool) bool {
	return true
}
