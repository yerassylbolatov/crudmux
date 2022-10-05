package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	database "github.com/yerassylbolatov/crudmux/db"
	"github.com/yerassylbolatov/crudmux/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error at initConfig occured %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env parameters: %s", err.Error())
	}
	db, err := database.DbStart(&database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		User:     viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	handlers.Db = db
	routes := handlers.InitRoutes()
	fmt.Println("server starts at :8080")
	log.Fatal(http.ListenAndServe(viper.GetString("port"), routes))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
