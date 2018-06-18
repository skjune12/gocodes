package main

import (
	"log"
	"net"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var ips []string
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range ifaces {
		// interface down
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		// loopback interface
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			log.Fatal(err)
		}

		for _, addr := range addrs {
			ips = append(ips, addr.String())
		}
	}
	spew.Dump(ips)
}
