// Server for RFC 868
// Usage: run server -p 11037
package main

import (
	"log"
	"os"
	"timeServer/server/srv"
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

	// run server
	srv.RunServer(port)
}
