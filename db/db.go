package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//const (
//	Host    = "localhost"
//	User    = "postgres"
//	DabName = "postgres"
//)

type Config struct {
	Host     string
	Port     int
	User     string
	DbName   string
	Password string
}

func DbStart(config *Config) (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
