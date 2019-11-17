package main

import (
	"fmt"

	"auth/server"
)

func main() {
	s := server.NewServer(8080)
	s.Listen()
	fmt.Println("Terminating Server")
}
