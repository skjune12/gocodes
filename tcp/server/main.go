package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)

	ln, err := net.Listen("tcp", ":7")
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accept new connection")
		go echo(con)
	}
}

func echo(con net.Conn) {
	_, err := io.Copy(con, con)
	if err != nil {
		log.Print(err)
	}

	err = con.Close()
	if err != nil {
		log.Println(err)
	}
}
