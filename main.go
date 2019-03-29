package main

import (
	"log"
	"net"
	"os"
)

func main() {

	_, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}
