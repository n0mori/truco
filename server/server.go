package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

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

func play(control *connectionControl, gameState *truco.GameState) {

	for gameState.PlayerStates[0].Tentos < 12 && gameState.PlayerStates[1].Tentos < 12 {
		playHand(control, gameState)
	}

	sendStates(control, gameState)

	closeConnections(control)
}

func playHand(control *connectionControl, gameState *truco.GameState) {
	var turn int
	if gameState.Round%2 == 0 {
		gameState.PlayerStates[0].Active = true
		gameState.PlayerStates[1].Active = false
		turn = 0
	} else {
		gameState.PlayerStates[0].Active = false
		gameState.PlayerStates[1].Active = true
		turn = 1
	}
	for gameState.PlayerStates[0].Maos < 3 && gameState.PlayerStates[1].Maos < 3 {
		values := make([]int, 2)

		play, err := control.readers[turn].ReadString('\n')
		if err == io.EOF {
			os.Exit(1)
		}

		ind, _ := strconv.Atoi(play)

		values[turn] = gameState.PlayerStates[turn].Cards[ind].Value()
		gameState.PlayerStates[turn].Cards = append(gameState.PlayerStates[turn].Cards[:ind], gameState.PlayerStates[turn].Cards[ind+1:]...)

		turn = turn + 1
		turn = turn % 2

		play, err = control.readers[turn].ReadString('\n')
		if err == io.EOF {
			os.Exit(1)
		}

		ind, _ = strconv.Atoi(play)

		values[turn] = gameState.PlayerStates[turn].Cards[ind].Value()
		gameState.PlayerStates[turn].Cards = append(gameState.PlayerStates[turn].Cards[:ind], gameState.PlayerStates[turn].Cards[ind+1:]...)

		if values[0] == values[1] {
			gameState.PlayerStates[0].Maos++
			gameState.PlayerStates[1].Maos++
		} else if values[0] > values[1] {
			gameState.PlayerStates[0].Maos++
		} else {
			gameState.PlayerStates[1].Maos++
		}

		sendStates(control, gameState)

	}

	if gameState.PlayerStates[0].Maos > gameState.PlayerStates[1].Maos {
		gameState.PlayerStates[0].Tentos++
	} else if gameState.PlayerStates[0].Maos < gameState.PlayerStates[1].Maos {
		gameState.PlayerStates[1].Tentos++
	}

	gameState.PlayerStates[0].Maos = 0
	gameState.PlayerStates[1].Maos = 0
	gameState.Round++
}

func closeConnections(control *connectionControl) {
	control.conns[0].Close()
	control.conns[1].Close()
}

func sendStates(control *connectionControl, gameState *truco.GameState) {
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

	play(&control, &gameState)
}
