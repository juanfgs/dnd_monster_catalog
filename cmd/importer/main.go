package main

import (
	"github.com/juanfgs/dnd-monster-library/internal/config"
	"github.com/juanfgs/dnd-monster-library/internal/db"
	"github.com/juanfgs/dnd-monster-library/internal/monster"
	"github.com/juanfgs/dnd-monster-library/internal/stats"
	"github.com/juanfgs/dnd-monster-library/internal/loader"
	"context"
	"log"
)


func main() {
	var config = config.NewConfig()
	monsters := loader.LoadMonsters()
	config.ReadEnv()
	pool, err := db.Connect(config.DSN)
	if err != nil {
		log.Fatal(err)
	}
	
	monsterRepo := monster.NewRepository(pool)

	statsRepo := stats.NewRepository(pool)
	ctx := context.Background()
	for _, m := range(monsters) {
		log.Println(m)
		err := monsterRepo.Create(ctx, &m)
		if err != nil {
			log.Fatal(err)
		}
		err = statsRepo.Create(ctx, m.Stats, m.ID)
		if err != nil {
			log.Fatal(err)
		}
	}

	
}

