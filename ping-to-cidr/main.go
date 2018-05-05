package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
)

// https://play.golang.org/p/m8TNTtygK0
func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// when CIDR is /31
	if len(ips) == 2 {
		return ips, nil
	}
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

type Pong struct {
	IP    string `json:"ip",omitempty`
	Alive bool   `json:"alive",omitempty`
}

func ping(addr string) Pong {
	_, err := exec.Command("ping", "-c1", "-t1", addr).Output()
	var alive bool
	if err != nil {
		alive = false
	} else {
		alive = true
	}
	return Pong{IP: addr, Alive: alive}
}

func main() {
	hosts, err := Hosts("203.178.143.0/24")
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range hosts {
		pong := ping(ip)
		fmt.Println(pong)
	}
}
