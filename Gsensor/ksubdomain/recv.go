package ksubdomain

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
	"sync/atomic"
	"time"
)

func Recv(device string, flagID uint16, retryChan chan RetryStruct) {
	var (
		snapshotLen int32         = 1024
		promiscuous bool          = false
		timeout     time.Duration = -1 * time.Second
	)
	//windowWith := GetWindowWith()

	handle, _ := pcap.OpenLive(device, snapshotLen, promiscuous, timeout)
	err := handle.SetBPFFilter("udp and src port 53")
	if err != nil {
		log.Fatalf("SetBPFFilter Faild:%s\n", err.Error())
	}
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	defer handle.Close()

	var udp layers.UDP
	var dns layers.DNS
	var eth layers.Ethernet
	var ipv4 layers.IPv4
	var ipv6 layers.IPv6
	//var subNextData []string

	parser := gopacket.NewDecodingLayerParser(
		layers.LayerTypeEthernet, &eth, &ipv4, &ipv6, &udp, &dns)

	for {
		packet, err := packetSource.NextPacket()
		if err != nil {
			continue
		}
		var decoded []gopacket.LayerType
		err = parser.DecodeLayers(packet.Data(), &decoded)
		if err != nil {
			continue
		}
		if !dns.QR {
			continue
		}
		ResultObj.AddCount()

		if dns.ID/100 == flagID {

			atomic.AddUint64(&RecvIndex, 1)
			udp, _ := packet.Layer(layers.LayerTypeUDP).(*layers.UDP)
			index := GenerateMapIndex(dns.ID%100, uint16(udp.DstPort))
			_, err := LocalStauts.SearchFromIndexAndDelete(uint32(index))
			if err == nil {
				if LocalStack.Len() <= 50000 {
					LocalStack.Push(uint32(index))
				}
			}
			if dns.ANCount > 0 {
				atomic.AddUint64(&SuccessIndex, 1)
				if len(dns.Questions) == 0 {
					continue
				}
				data := RecvResult{Subdomain: string(dns.Questions[0].Name)}
				data.Answers = dns.Answers

				msg := data.Subdomain + " => "

				msg = strings.Trim(msg, " => ")
				dnsinfo := ""
				for _, v := range data.Answers {
					dnsinfo += v.String()
					dnsinfo += ","
				}
				log.Println(msg) // result
				ResultObj.AddResult(data.Subdomain, dnsinfo)
				//ff := windowWith - len(msg) - 1
			}
		}
	}
}
