package main

import (
	"fmt"
	_ "github.com/lib/pq"
	database "github.com/yerassylbolatov/crudmux/db"
	"github.com/yerassylbolatov/crudmux/handlers"
	"log"
	"net/http"
)

func main() {
	db, err := database.DbStart()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	handlers.Db = db
	routes := handlers.InitRoutes()
	fmt.Println("server starts at :8080")
	log.Fatal(http.ListenAndServe(":8080", routes))
}
