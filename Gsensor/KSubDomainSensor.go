package Gsensor

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/code-scan/Goal/Gsensor/ksubdomain"
	rateLimit "golang.org/x/time/rate"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"
)

type KSubDomainSensor struct {
	Tasks  []string
	Target string
	Type   string
	result Result
}

func NewKSubDomainSensor() *KSubDomainSensor {
	s := KSubDomainSensor{}
	return &s
}

func (s *KSubDomainSensor) GetInfo() string {
	return "KSubDomainSensor ver 0.1 with  " + s.Type

}

func (s *KSubDomainSensor) SetDomain(domain string) {
	s.Target = domain
}

func (s *KSubDomainSensor) SetAccount(_ string) {
	return
}

func (s *KSubDomainSensor) SetPassword(_ string) {
	return
}

func (s *KSubDomainSensor) SetType(type_ string) {
	s.Type = type_
}

//go:embed subnames_test.txt
var data string

func (s *KSubDomainSensor) AddTask() {
	s.Tasks = []string{}
	datas := strings.Split(string(data), "\n")
	for _, sub := range datas {
		sub = strings.TrimSpace(sub)
		domain := fmt.Sprintf("%s.%s", sub, s.Target)
		s.Tasks = append(s.Tasks, domain)
	}
}
func (s *KSubDomainSensor) GetResult() Result {
	s.result = Result{}
	if s.Type != "subdomain" {
		return s.result
	}
	s.AddTask()
	ksubdomain.LocalStack = ksubdomain.NewStack()
	ksubdomain.ResultObj.Length = len(s.Tasks)
	var rate int64
	rate = 1000
	// 获取网卡
	ether := ksubdomain.AutoGetDevices()
	// packet 标识符
	flagID := uint16(ksubdomain.RandInt64(400, 654))
	// 用于重试的chan
	retryChan := make(chan ksubdomain.RetryStruct, rate)
	go ksubdomain.Recv(ether.Device, flagID, retryChan)
	sendog := ksubdomain.SendDog{}
	sendog.Init(ether, []string{"223.5.5.5", "223.6.6.6", "180.76.76.76", "119.29.29.29", "178.79.131.110", "199.91.73.222", "185.184.222.222", "185.222.222.222", "182.254.116.116", "140.207.198.6", "123.125.81.6", "114.114.114.115", "218.30.118.6", "101.226.4.6"}, flagID, true)
	limiter := rateLimit.NewLimiter(rateLimit.Every(time.Duration(time.Second.Nanoseconds()/rate)), int(rate))
	ctx := context.Background()

	// 协程重发检测
	go func() {
		for {
			// 循环检测超时的队列
			maxLength := int(rate / 10)
			datas := ksubdomain.LocalStauts.GetTimeoutData(maxLength)
			isdelay := true
			if len(datas) <= 100 {
				isdelay = false
			}
			for _, localdata := range datas {
				index := localdata.Index
				value := localdata.V
				if value.Retry >= 15 {
					atomic.AddUint64(&ksubdomain.FaildIndex, 1)
					ksubdomain.LocalStauts.SearchFromIndexAndDelete(index)
					continue
				}
				_ = limiter.Wait(ctx)
				value.Retry++
				value.Time = time.Now().Unix()
				value.Dns = sendog.ChoseDns()
				// 先删除，再重新创建
				ksubdomain.LocalStauts.SearchFromIndexAndDelete(index)
				ksubdomain.LocalStauts.Append(&value, index)
				flag2, srcport := ksubdomain.GenerateFlagIndexFromMap(index)
				retryChan <- ksubdomain.RetryStruct{Domain: value.Domain, Dns: value.Dns, SrcPort: srcport, FlagId: flag2, DomainLevel: value.DomainLevel}
				if isdelay {
					time.Sleep(time.Microsecond * time.Duration(rand.Intn(300)+100))
				}
			}
		}
	}()
	// 多级域名检测
	go func() {
		for {
			rstruct := <-retryChan
			if rstruct.SrcPort == 0 && rstruct.FlagId == 0 {
				flagid2, scrport := sendog.BuildStatusTable(rstruct.Domain, rstruct.Dns, rstruct.DomainLevel)
				rstruct.FlagId = flagid2
				rstruct.SrcPort = scrport
			}
			_ = limiter.Wait(ctx)
			sendog.Send(rstruct.Domain, rstruct.Dns, rstruct.SrcPort, rstruct.FlagId)
		}
	}()
	for _, msg := range s.Tasks {
		dnsname := sendog.ChoseDns()
		flagid2, scrport := sendog.BuildStatusTable(msg, dnsname, 1)
		sendog.Send(msg, dnsname, scrport, flagid2)
	}
	for {
		if ksubdomain.LocalStauts.Empty() {
			break
		}
		time.Sleep(time.Millisecond * 723)
	}
	s.result = ksubdomain.ResultObj.GetResult()
	return s.result
}
