package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal()
	}

	defer li.Close()

	conn, err := li.Accept()
	if err != nil {
		log.Fatal()
	}

	defer conn.Close()

	io.WriteString(conn, "I see you connected")
}
