package config

import (
	"os"
	"fmt"
)

type Config struct {
	DSN string
}

func NewConfig() Config {
	var config Config  
	return config

}

func (c *Config) ReadEnv() *Config {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	c.DSN = buildDSN(user, pass, host, port, dbname)
	return c
}

func buildDSN(user string, pass string, host string, port string, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
}

