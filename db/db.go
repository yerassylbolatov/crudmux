package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

const (
	Host    = "localhost"
	Port    = "5432"
	User    = "postgres"
	DabName = "postgres"
)

func StartDb() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error when loading environment variables %s\n", err.Error())
		return
	}
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, os.Getenv("DB_PASSWORD"), DabName)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Printf("error when openinig sql %s\n", err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("error when pinging sql %s\n", err.Error())
		return
	}
	fmt.Println("db connected successfully!")
}
