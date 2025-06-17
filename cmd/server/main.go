package main

import (
	"github.com/juanfgs/dnd-monster-library/internal/server"
	"github.com/juanfgs/dnd-monster-library/internal/config"
	"github.com/juanfgs/dnd-monster-library/internal/db"
	"log"
)
 

func  main(){
	var config = config.NewConfig()
	config.ReadEnv()

	pool, err := db.Connect(config.DSN)
	if err != nil {
		log.Fatal(err)
	}
	
	s := server.NewServer(config, pool)
	s.Listen(config.HTTP.Port)
}
