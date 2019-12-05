package main

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"os"
	"time"
)

func ping(proto, ip string) {
	p := fastping.NewPinger()
	//ra, err := net.ResolveIPAddr("ip4:icmp", os.Args[1])
	ra, err := net.ResolveIPAddr(proto, ip)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ping("ip6:icmp", "::1")
	ping("ip4:icmp", "192.168.6.55")
	ping("ip4:icmp", "192.168.6.195")
	ping("ip6:icmp", "2001:4860:4860::8888")
}
