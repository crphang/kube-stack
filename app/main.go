package main

import (
	"fmt"

	"app/server"
)

func main() {
	// http.HandleFunc("/", handler)
	s := server.NewServer(3000)
	s.Listen()
	fmt.Println("Terminating Server")
}
