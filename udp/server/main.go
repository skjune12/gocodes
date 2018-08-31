package main

import (
	"log"
	"net"
	"os"
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	udpLn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	log.Println("Starting udp server")

	for {
		n, addr, err := udpLn.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}

		go func() {
			log.Printf("Receiving data: %s from %s\n", string(buf[:n]), addr.String())
			log.Println("Sending data...")
			udpLn.WriteTo([]byte("Pong"), addr)
			log.Printf("Complete Sending Data")
		}()
	}
}
