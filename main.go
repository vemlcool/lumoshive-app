package main

import (
	"log"
	"myproject/database"
	"myproject/routes"
	"net/http"
)

func main() {
	database.ConnectDatabase()
	mux := routes.SetAPI()
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
