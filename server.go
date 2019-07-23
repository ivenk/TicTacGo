package main

import (
	"./codes"
	"./game"
	"fmt"
	"net"
	"os"
)

const (
	PORT            = "8765"
	HOST            = "localhost"
	CONNECTION_TYPE = "tcp"
)

type Player struct {
	number int
	conn   net.Conn
}

func main() {
	//set up listener
	listener, err := net.Listen(CONNECTION_TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Printf("Error while listening on port: %s and host: %s    : %s", PORT, HOST, err.Error())
		os.Exit(1)
	}
	// partly taken from https://pocketgophers.com/handling-errors-in-defer/
	// allows error handling with defer
	defer func() {
		err := listener.Close()
		if err != nil {
			fmt.Printf("Error while shutting down listener on : %s:%s    : %s", PORT, HOST, err.Error())
			os.Exit(1)
		}
	}()

	// initialise game
	game := game.Ticktacktoe{}
	c := make(chan int)

	var players []Player
	i := 0
	for len(players) != 2 {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection : %s", err.Error())
			os.Exit(1)
		}
		players = append(players, Player{i, conn})
		i = i + 1

		go handleRequest(game, players[(i-1)], c)
	}
}

func handleRequest(game game.Ticktacktoe, player Player, c chan int) {
	defer func() {
		err := player.conn.Close()
		if err != nil {
			fmt.Printf("Error while closing connection on %s:%s      : %s", HOST, PORT, err.Error())
		}
	}()

	if player.number == 0 {
		player.conn.Write([]byte(string(codes.WAITING_FOR_PLAYERS)))
		signal := <-c
		if signal != codes.GAME_STARTING {
			// error!
			fmt.Println("Wrong signal received during game setup")
			os.Exit(0)
		}
	}

	if player.number == 1 {
		c <- codes.GAME_STARTING
	}

	player.conn.Write([]byte(string(codes.GAME_STARTING)))
	for {
		buf := make([]byte, 1024)
		length, err := player.conn.Read(buf)
		// log but dont crash server
		if err != nil {
			fmt.Printf("Error receiving message on: %s:%s      : %s", HOST, PORT, err.Error())
		}

		fmt.Printf("Received message: %s", buf[:length])
	}
}
