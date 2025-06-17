package server


import (
	"net/http"
	"github.com/juanfgs/dnd-monster-library/internal/config"
	"log"
	"database/sql"
)


type Server struct {
	db *sql.DB
	mux *http.ServeMux
}

func NewServer(c config.Config, pool *sql.DB ) *Server {
	if c.HTTP == nil {
		log.Fatal("Unable to load HTTP Server configuration.")
	}

	s := &Server{
		db: pool,
		mux:  http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) Listen(addr string) error {
    return http.ListenAndServe(addr, s.mux)
}
