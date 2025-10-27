package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/users", CreateUserHandler)

	port := 8080
	log.Printf("🚀 Server läuft auf http://localhost:%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("❌ Serverfehler: %v", err)
	}
}
