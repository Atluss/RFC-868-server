// Server for RFC 868
// Usage: run server -p 11037
package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"timeServer/pkg/v1"
)

func main() {
	var port string
	var err error
	// validate program arguments
	if port, err = v1.CheckServerSettings(os.Args[1:]); err != nil {
		log.Println(err)
		return
	}
	// run server
	runServer(port)
}

func runServer(port string) {

	address := fmt.Sprintf("%s:%s", v1.ConnHost, port)

	l, err := net.Listen(v1.ConnType, address)
	if err != nil {
		fmt.Printf("error: listening: %s", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	log.Printf("listening on %s:%s", v1.ConnHost, port)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()

		log.Printf("request from: %s", conn.RemoteAddr())

		if err != nil {
			log.Printf("error accepting: %s", err.Error())
			os.Exit(1)
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {

	readBuffer := make([]byte, 1024)
	buf := make([]byte, 4)

	_, err := conn.Read(readBuffer)

	if err != nil {
		log.Println("error reading:", err.Error())
	}

	binary.BigEndian.PutUint32(buf, v1.RFC868Time())

	if _, err := conn.Write(buf); err != nil {
		log.Println("error: can't write respond")
	}

	if err := conn.Close(); err != nil {
		log.Println("error: server can't close connection")
	}
}
