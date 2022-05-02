package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Launching server...")

	// Setting up port listening
	ln, err := net.Listen("tcp", ":8081")
	logFatal(err)

	// Opening the port
	conn, err := ln.Accept()
	logFatal(err)

	// listening to all messages separated by \n
	message, err := bufio.NewReader(conn).ReadString('\n')
	logFatal(err)

	// Printing received message
	fmt.Print("Message Received:", string(message))

	// Sending a new line back to the client
	conn.Write([]byte(message))
}

func logFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
