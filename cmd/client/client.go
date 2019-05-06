package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"timeServer/pkg/v1"
)

func main() {

	var url, port string
	var err error

	// validate program arguments
	if url, port, err = v1.CheckClientSettings(os.Args[1:]); err != nil {
		log.Println(err)
		return
	}

	address := fmt.Sprintf("%s:%s", url, port)

	log.Printf("Address to connect: %s", address)
	for {
		if str, err := v1.DialToTimeServer(address); err != nil {
			log.Println(err)
			return
		} else {
			log.Println(str)
		}
		time.Sleep(time.Second)
	}

}
