package config

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"log"
)

type HTTP struct {
	Port string
	ReadTimeout time.Duration 
	WriteTimeout time.Duration
}

type Config struct {
	DSN string
	HTTP *HTTP
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
	readTimeout := getIntOrFail("HTTP_READ_TIMEOUT")
	writeTimeout := getIntOrFail("HTTP_WRITE_TIMEOUT")
	httpConf := &HTTP{
		Port: os.Getenv("HTTP_PORT"),
		ReadTimeout: time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	}

	c.DSN = buildDSN(user, pass, host, port, dbname)
	c.HTTP = httpConf
	return c
}

func getIntOrFail(key string) int64 {
	value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		log.Fatal("Invalid value for %s", key )
	}
	return value
}

func buildDSN(user string, pass string, host string, port string, dbname string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbname)
}

