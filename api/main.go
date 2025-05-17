package main

import (
	"log"
	"net/http"
	"project/config"
	"project/routes"
)

func main() {
	config.ConnectDB()
	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
