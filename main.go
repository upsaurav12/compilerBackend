package main

import (
	"log"
	"net/http"
	"online-compiler/handlers"
)

func main() {
	http.HandleFunc("/execute", handlers.HandleExecute)

	log.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
