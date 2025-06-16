package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"context"
	"time"
	"log"
)

var pool *sql.DB

func Connect(dsn string) (*sql.DB, error) {
	var err error 
	pool, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Println("unable to use data source name", err  )
		return nil, err
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()


	Ping(ctx)
	return pool, nil 
}

// Ping the database to verify DSN provided by the user is valid and the
// server accessible. If the ping fails exit the program with an error.
func Ping(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}
