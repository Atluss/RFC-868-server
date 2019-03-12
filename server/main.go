// Server for RFC 868
// Usage: run server -p 11037
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"timeServer/utl"
)


const (
	ConnHost = "localhost"
	ConnType = "tcp"
)

func main() {

	var port string
	var err error

	// validate program arguments
	if port, err = utl.CheckServerSettings(os.Args[1:]); err != nil {
		log.Println(err)
		return
	}

	l, err := net.Listen(ConnType, ConnHost+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()
	log.Println("Listening on " + ConnHost + ":" + port)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()

		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	//buf := make([]byte, 1024)
	buf := make([]byte, 4)
	_, err := conn.Read(buf)

	if err != nil {
		log.Println("Error reading:", err.Error())
	}

	binary.BigEndian.PutUint32(buf, utl.RFC868Time())

	conn.Write(buf)

	conn.Close()
}

