package encounter

import (
	"database/sql"
	"encoding/json"
	"io"
	"context"
	"net/http"
	"github.com/juanfgs/dnd-monster-library/internal/monster"
)
func  CreateHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
	    monsters := make([]monster.Monster, 0 ,0 )
	    repo := monster.NewRepository(db)
	    ctx := context.Background()
	    if r.Method != http.MethodPost {
		    http.Error(w,"Method not Allowed", http.StatusMethodNotAllowed)
		    return 
	    }

	    body ,err := io.ReadAll(r.Body)
	    if err != nil {
		    http.Error(w,"Unable to read body", http.StatusBadRequest)
		    return 
	    }

	    defer r.Body.Close()
	    
	    var request CreateRequest
	    err = json.Unmarshal(body, &request)

	    if err != nil {
		    http.Error(w,"Error decoding JSON", http.StatusBadRequest)
		    return
	    }

	    if request.Quantity <= 0 {
		    http.Error(w,"Quantity must be higher than zero ", http.StatusBadRequest)
		    return

	    }

	    monsters, err  = repo.FindByChallengeRating(ctx,
		    request.MinChallengeRating,
		    request.MaxChallengeRating,
		    request.Quantity)
	    w.Header().Set("Content-Type", "application/json")
	    json.NewEncoder(w).Encode(monsters)
    }
} 

