package server

import (
	"fmt"
	"log"
	"net/http"

	"app/server/api"
)

// Server interface
type Server interface {
	Listen()
}

// SimpleHTTPServer implements Server Implmentation
type SimpleHTTPServer struct {
	Port int
}

// NewServer initializes a SimpleServer
func NewServer(port int) *SimpleHTTPServer {
	s := &SimpleHTTPServer{
		Port: port,
	}

	s.routes()
	return s
}

// Listen starts listening
func (s *SimpleHTTPServer) Listen() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil))
}

func (s *SimpleHTTPServer) routes() {
	http.HandleFunc("/", api.Index)
}
