package main 

import (
	"github.com/juanfgs/dnd-monster-library/internal/config"
	"github.com/juanfgs/dnd-monster-library/internal/db"
)

func main() {
	var config = config.NewConfig()
	config.ReadEnv()
	db.Connect(config.DSN)
}
