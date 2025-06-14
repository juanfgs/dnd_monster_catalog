package importer

import (
	"encoding/json" 
	"fmt" 
	"bytes" 
	"os" 
	"io" 
	"log" 
)

/*
 * Parses the monster file and loads all entities
 */
func ParseFile() {
	file, err := os.ReadFile("./data/5e-SRD-Monsters.json")

	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(bytes.NewReader(file))
	for {
		var monsters []Monster 
		if err := dec.Decode(&monsters); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for _, m := range monsters {
			log.Println(m)
			fmt.Printf("%s: %s\n", m.Name, m.Size)
		}
	}
}



