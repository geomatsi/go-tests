// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{}

func control(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer conn.Close()

	cli := make(chan string)
	net := make(chan string)

	go handleCli(cli)
	go handleNet(net, conn)

	for {
		select {
		case message, more := <-cli:
			if more == false {
				return
			}

			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				return
			}

		case message, more := <-net:
			if more == false {
				return
			}

			fmt.Printf("recv: %s\n", message)
		}
	}
}

func handleCli(cc chan string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("err: stdio read: %s\n", err)
			}

			close(cc)
			return
		}

		cc <- str
	}
}

func handleNet(cc chan string, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			close(cc)
			return
		}

		cc <- string(message)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "websocket test")
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/control", control)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
