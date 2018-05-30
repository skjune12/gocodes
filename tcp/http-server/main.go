package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// 一度で終了しないため無限ループ
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				log.Fatal(err)
			}

			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(dump))

			// レスポンスを書き込む
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       ioutil.NopCloser(strings.NewReader("Hello, World!")),
			}

			response.Write(conn)
			conn.Close()
		}()
	}
}
