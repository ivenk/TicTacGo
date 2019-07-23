package main

import (
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

	// main loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection : %s", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn, game)
	}
}

func handleRequest(conn net.Conn, game game.Ticktacktoe) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error while closing connection on %s:%s      : %s", HOST, PORT, err.Error())
		}
	}()

	buf := make([]byte, 1024)
	length, err := conn.Read(buf)
	// log but dont crash server
	if err != nil {
		fmt.Printf("Error receiving message on: %s:%s      : %s", HOST, PORT, err.Error())
	}

	fmt.Printf("Message received from %s : %s", conn.RemoteAddr(), buf[:length])
}
