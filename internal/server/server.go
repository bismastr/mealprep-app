package server

import (
	"fmt"
	"net/http"

	"github.com/bismastr/mealprep-app/internal/db"
)

type Server struct {
	server *http.Server
	db     *db.DB
}

func NewServer(db *db.DB) *Server {
	return &Server{
		server: &http.Server{
			Addr: ":8084",
		},
		db: db,
	}
}

func (s *Server) Start() error {
	fmt.Println("Server listening on port 8084")
	s.server.Handler = s.RegisterRoute()
	return s.server.ListenAndServe()
}
