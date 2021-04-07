package Gnet

import (
	"context"
	"net"
	"time"
)

func GetHostIp(domain string, dns string) []string {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: 10 * time.Second,
			}
			return d.DialContext(ctx, "udp", dns)
		},
	}

	ips, _ := r.LookupHost(context.Background(), domain)
	return ips
}
