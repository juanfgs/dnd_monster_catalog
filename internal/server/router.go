package server

import (
	"github.com/juanfgs/dnd-monster-library/internal/monster"
)

func (s *Server) routes() {
	s.mux.HandleFunc("/", monster.ListHandler(s.db))
}
