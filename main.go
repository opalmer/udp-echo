package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	listen = flag.String("listen", "", "Listen on the given addr:port combination.")
	buffer = flag.Int("buffer", 1024, "How much data to buffer.")
	noout  = flag.Bool("no-output", false, "If provided then only print errors")
)

func main() {
	flag.Parse()
	if len(*listen) == 0 {
		log.Fatal("-listen not provided")
	}

	addr, err := net.ResolveUDPAddr("udp", *listen)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening...")

	server, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		buf := make([]byte, *buffer)
		size, _, err := server.ReadFromUDP(buf)
		if err != nil {
			log.Println("ERROR: ", err)
			continue
		}
		if !*noout {
			log.Println(string(buf[0:size]))
		}
	}
}
