package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/n0mori/truco/lib"
)

type connectionControl struct {
	conns   []*net.TCPConn
	readers []*bufio.Reader
}

func waitForPlayers(server *net.TCPListener) connectionControl {
	conns := make([]*net.TCPConn, 2)
	readers := make([]*bufio.Reader, 2)

	number := 0
	for number < 2 {
		conn, err := server.AcceptTCP()

		if err != nil {
			log.Fatal(err)
		}

		println(conn.RemoteAddr())

		reader := bufio.NewReader(conn)

		conns[number] = conn
		readers[number] = reader

		fmt.Fprintln(conn, number)
		number++
	}
	return connectionControl{conns: conns, readers: readers}
}

func play(control connectionControl, gameState truco.GameState) {
	sendStates(control, gameState)
}

func sendStates(control connectionControl, gameState truco.GameState) {
	for _, conn := range control.conns {
		js, err := json.Marshal(gameState)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(conn, string(js))
		fmt.Println(string(js))
	}

}

func main() {
	tcp, err := net.ResolveTCPAddr("tcp", ":2000")

	if err != nil {
		log.Fatal(err)
	}

	server, err := net.ListenTCP("tcp", tcp)

	if err != nil {
		log.Fatal(err)
	}

	control := waitForPlayers(server)

	gameState := truco.StartGame()

	play(control, gameState)
}
