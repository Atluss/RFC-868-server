package srv

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"timeServer/utl"
)

func RunServer(port string) {

	address := fmt.Sprintf("%s:%s", utl.ConnHost, port)

	l, err := net.Listen(utl.ConnType, address)
	if err != nil {
		fmt.Printf("error: listening: %s", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	log.Printf("listening on %s:%s", utl.ConnHost, port)

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

	binary.BigEndian.PutUint32(buf, utl.RFC868Time())

	conn.Write(buf)

	conn.Close()
}
