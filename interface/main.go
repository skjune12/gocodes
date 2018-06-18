package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var devname = flag.String("dev", "", "name of interface")

func init() {
	flag.Parse()
}

func main() {
	var ips []net.IP

	reader := bufio.NewReader(os.Stdin)

	if *devname == "" {
		fmt.Printf("Enter Device Name: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSuffix(input, "\n")
		devname = &input
	}

	dev, err := net.InterfaceByName(*devname)
	if err != nil {
		log.Fatal(err)
	}

	addrs, err := dev.Addrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			log.Fatal(err)
		}

		if ip.IsGlobalUnicast() {
			ips = append(ips, ip)
		}
	}

	spew.Dump(ips)
}
