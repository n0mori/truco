package truco

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

//Player is a struct that creates the connection and gamerules
type Player struct {
	server   *net.TCPConn
	nextMove string
	active   bool
}

//NewPlayer returns a new instance of Player
func NewPlayer(ip string) Player {
	tcp, err := net.ResolveTCPAddr("tcp", ip+":2000")

	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcp)

	if err != nil {
		log.Fatal(err)
	}

	return Player{server: conn, active: false}
}

//Move makes a move when it is possible
func (p Player) Move(str string) {
	if p.active != true {
		p.wait()
	}

	fmt.Fprintln(p.server, str)

}

func (p Player) wait() {
	reader := bufio.NewReader(p.server)

	for !p.active {
		str, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.HasPrefix(str, "turn") {
			break
		}
	}
}

//Close closes the connection
func (p Player) Close() {
	p.server.Close()
}
