package main

import (
	"log"
	"net/http"

	"github.com/breadwithmeth/naliv_go/handlers"
)

func main() {
	http.HandleFunc("/items", handlers.ItemsHandler)

	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
