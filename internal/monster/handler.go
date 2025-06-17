package monster

import (
	"net/http"
	"database/sql"
)


func  ListHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
	    
	    repo := NewRepository(db)
        monsters, err := repo.Index(db)
        if err != nil {
            http.Error(w, "error fetching monsters", 500)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(monsters)
    }
} 
