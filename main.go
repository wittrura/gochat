package main

import (
	"flag" // abstraction of os.Args
	"gochat/lib"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified ip address")
	flag.Parse()

	if isHost {
		connIP := os.Args[2]
		lib.RunHost(connIP)
	} else {
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}

}
