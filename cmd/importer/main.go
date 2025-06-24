package main

import (
	"context"
	"log"

	"github.com/juanfgs/dnd-monster-library/internal/armor"
	"github.com/juanfgs/dnd-monster-library/internal/config"
	"github.com/juanfgs/dnd-monster-library/internal/db"
	"github.com/juanfgs/dnd-monster-library/internal/loader"
	"github.com/juanfgs/dnd-monster-library/internal/monster"
	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
	"github.com/juanfgs/dnd-monster-library/internal/speed"
	"github.com/juanfgs/dnd-monster-library/internal/stats"
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
	proficiencyRepo := proficiency.NewRepository(pool)
	armorRepo := armor.NewRepository(pool)
	speedRepo := speed.NewRepository(pool) 
	
	ctx := context.Background()
	for _, mDTO := range(monsters) {
		m := mDTO.BuildModel()
		err := monsterRepo.Create(ctx, m)
		if err != nil {
			log.Fatal(err)
		}
		err = statsRepo.Create(ctx, m.Stats, m.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, pDTO := range(mDTO.Proficiencies) {
			p := pDTO.BuildModel()
			err = proficiencyRepo.Create(ctx, &p)
			if err != nil {
				log.Fatal(err)
			}
		        err = proficiencyRepo.Associate(ctx, p.ID, m.ID, pDTO.Value)	
			if err != nil {
				log.Fatal(err)
			}
		}
		for _, acDTO := range(mDTO.ArmorClass) {
			ac := acDTO.BuildModel()
			err = armorRepo.Create(ctx, &ac)
			if err != nil {
				log.Fatal(err)
			}
		        err = armorRepo.Associate(ctx, ac.ID, m.ID, acDTO.Value)	
			if err != nil {
				log.Fatal(err)
			}
		}
		for _, s := range( speed.BuildModels(mDTO.Speed)) {
			err = speedRepo.Create(ctx, &s)
			if err != nil {
				log.Fatal(err)
			}
			err = speedRepo.Associate(ctx, s.ID, m.ID, s.Value, s.Unit)	
			if err != nil {
				log.Fatal(err)
			}
		}
		

	}

	
}

