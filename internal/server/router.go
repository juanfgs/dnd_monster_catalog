package server

import (
	"github.com/juanfgs/dnd-monster-library/internal/monster"
	"github.com/juanfgs/dnd-monster-library/internal/encounter"
)

func (s *Server) routes() {
	s.mux.HandleFunc("/", monster.ListHandler(s.db))
	s.mux.HandleFunc("/encounter", encounter.CreateHandler(s.db))
}
