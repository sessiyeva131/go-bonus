package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Connecting to the socket
	conn, err := net.Dial("tcp", "localhost:8081")
	logFatal(err)

	// Reading input data from STDIN
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? \n")
	text, err := reader.ReadString('\n')
	logFatal(err)

	// Sending to socket
	fmt.Fprintf(conn, text)

	// Listening to the response
	newMessage, err := bufio.NewReader(conn).ReadString('\n')
	logFatal(err)

	msg := fmt.Sprintf("Hello, %s", newMessage)
	fmt.Println(msg)

}

func logFatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
