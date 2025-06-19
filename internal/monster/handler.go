package monster

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"context"
	"log"
)


func  ListHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
	    
	    ctx := context.Background()
	    repo := NewRepository(db)
	    monsters, err := repo.Index(ctx)
	    if err != nil {
		    log.Println(err)
		    http.Error(w, "error fetching monsters", 500)
		    return
	    }

	    w.Header().Set("Content-Type", "application/json")
	    json.NewEncoder(w).Encode(monsters)
    }
} 
