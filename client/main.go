package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"
	"timeServer/utl"
)

func main() {

	var port string
	var err error

	// validate program arguments
	if port, err = utl.CheckServerSettings(os.Args[1:]); err != nil {
		log.Println(err)
		return
	}

	address := fmt.Sprintf("%s:%s", utl.ConnHost, port)

	log.Printf("Address to connect: %s", address)
	for {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("error to connect: %s", err)
		}

		if _, err := fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n"); err != nil {
			log.Printf("error: %s", err)
		}

		if status, _, err := bufio.NewReader(conn).ReadLine(); err != nil {
			log.Println(err)
		} else {
			log.Println("respond: ", binary.BigEndian.Uint32(status))
		}

		time.Sleep(time.Second)
	}

}
