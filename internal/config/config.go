package config

import (
	"os"
	"fmt"
)

type Config struct {
	DSN string
}

func NewConfig() : Config {
	var config Config  
	return config

}

func (c *Config) ReadEnv() *Config {

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	buildDSN(user, pass, host, port, dbname)
}

func buildDSN(user string, pass string, host string, port string, dbname string) string {
	fmt.Sprintf("%s:%s@%s:%s/%s", user, pass, host, port, dbname)
}

