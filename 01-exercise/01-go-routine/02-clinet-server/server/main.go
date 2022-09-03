package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listner, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listner.Accept()

		if err != nil {
			log.Fatal(err)
			continue
		}

		go handleConn(conn)

	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, "Reponse fom Server")

		if err != nil {
			log.Fatal(err)
			return
		}

		time.Sleep(time.Second)
	}

}
