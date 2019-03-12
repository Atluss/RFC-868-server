package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"timeServer/utl"
)

func main() {

	var url, port string
	var err error

	// validate program arguments
	if url, port, err = utl.CheckClientSettings(os.Args[1:]); err != nil {
		log.Println(err)
		return
	}

	address := fmt.Sprintf("%s:%s", url, port)

	log.Printf("Address to connect: %s", address)
	for {
		utl.DialToTimeServer(address)
		time.Sleep(time.Second)
	}

}
