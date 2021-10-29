package ksubdomain

import (
	"log"
	"sync"
)

type Result struct {
	lock   sync.Mutex
	result map[string]string // url:code
	pos    int               // scan pos
	Length int               // ports length
}

func (d *Result) AddResult(subdomain string, dnsinfo string) {
	d.lock.Lock()
	d.result[subdomain] = dnsinfo
	d.lock.Unlock()
}
func (d *Result) GetPos() float64 {
	if d.Length == 0 {
		return 1
	}
	var t = float64(RecvIndex) / float64(d.Length)
	log.Println("getPos: ", t)
	return t
}
func (d *Result) AddCount() {
	d.lock.Lock()
	d.pos++
	d.lock.Unlock()
}
func (d *Result) GetResult() map[string]string {
	return d.result
}

var ResultObj = Result{
	result: make(map[string]string),
}
