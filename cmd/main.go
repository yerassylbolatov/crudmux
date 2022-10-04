package main

import (
	"fmt"
	"github.com/yerassylbolatov/crudmux/db"
	"github.com/yerassylbolatov/crudmux/handlers"
	"log"
	"net/http"
)

func main() {
	router := handlers.InitRoutes()
	db.StartDb()
	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("server starts at :8080")
}
