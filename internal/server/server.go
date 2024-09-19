package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer() *Server {
	return &Server{
		server: &http.Server{
			Addr: ":8080",
		},
	}
}

func (s *Server) Start() error {
	fmt.Println("Server listening on port 8080")
	s.server.Handler = RegisterRoute()
	return s.server.ListenAndServe()
}
