package main

import (
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
			fmt.Printf("Error while closing connection on : %s:%s    : %s", PORT, HOST, err.Error())
			os.Exit(1)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection : %s", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

}