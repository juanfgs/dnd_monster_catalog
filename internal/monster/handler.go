package monster

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/juanfgs/dnd-monster-library/internal/proficiency"
)


func  ListHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
	    
	    ctx := context.Background()
	    repo := NewRepository(db)
	    proficiencyRepo := proficiency.NewRepository(db)
	    monsters, err := repo.Index(ctx)
	    for i, m := range(monsters) {
		    var proficiencies []proficiency.Proficiency
		    proficiencies, err = proficiencyRepo.Fetch(ctx, m.ID)
		    monsters[i].Proficiencies = proficiencies
	    }
	    if err != nil {
		    log.Println(err)
		    http.Error(w, "error fetching monsters", 500)
	    }

	    w.Header().Set("Content-Type", "application/json")
	    json.NewEncoder(w).Encode(monsters)
    }
} 
