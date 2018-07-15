package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	number := 0

	tcp, err := net.ResolveTCPAddr("tcp", ":2000")

	if err != nil {
		log.Fatal(err)
	}

	server, err := net.ListenTCP("tcp", tcp)

	if err != nil {
		log.Fatal(err)
	}

	conns := make([]*net.TCPConn, 2)
	readers := make([]*bufio.Reader, 2)

	for number < 2 {
		conn, err := server.AcceptTCP()

		if err != nil {
			log.Fatal(err)
		}

		println(conn.RemoteAddr())

		reader := bufio.NewReader(conn)

		conns[number] = conn
		readers[number] = reader

		number++
	}

	number = 0
	for {
		println("wait")
		str := "turn player"
		fmt.Fprintln(conns[number%2], str)
		move, _ := readers[number%2].ReadString('\n')
		println(move)
		number++
	}
}
