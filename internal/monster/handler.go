package monster

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/juanfgs/dnd-monster-library/internal/armor"
	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
	"github.com/juanfgs/dnd-monster-library/internal/speed"
)


func  ListHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
	    
	    ctx := context.Background()
	    repo := NewRepository(db)
	    proficiencyRepo := proficiency.NewRepository(db)
	    armorRepo := armor.NewRepository(db)
	    speedRepo := speed.NewRepository(db) 
	    monsters, err := repo.Index(ctx)
	    for i, m := range(monsters) {
		    var proficiencies []proficiency.Proficiency
		    proficiencies, err = proficiencyRepo.Fetch(ctx, m.ID)
		    monsters[i].Proficiencies = proficiencies
	    }

	    for i, m := range(monsters) {
		    var armorClasses []armor.ArmorClass
		    armorClasses, err = armorRepo.Fetch(ctx, m.ID)
		    monsters[i].ArmorClasses = armorClasses 
	    }
	    for i, m := range(monsters) {
		    var speeds []speed.Speed
		    speeds, err = speedRepo.Fetch(ctx, m.ID)
		    monsters[i].Speed = speeds 
	    }
	    if err != nil {
		    log.Println(err)
		    http.Error(w, "error fetching monsters", 500)
	    }

	    w.Header().Set("Content-Type", "application/json")
	    json.NewEncoder(w).Encode(monsters)
    }
} 
