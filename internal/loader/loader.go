package loader

import (
	"encoding/json" 
	"github.com/juanfgs/dnd-monster-library/internal/monster"
	"bytes" 
	"os" 
	"io" 
	"log" 
)

/*
 * Parses the monster file and loads all entities
 */
func LoadMonsters() []monster.MonsterDTO {
	file, err := os.ReadFile("./data/5e-SRD-Monsters.json")

	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(bytes.NewReader(file))

	var monsters []monster.MonsterDTO
	for {
		if err := dec.Decode(&monsters); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		
	}
	
	return monsters
}
