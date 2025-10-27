package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Farben definieren (ANSI Terminalfarben)
const (
	green = "\033[32m" // ✅ Grün: Erfolg
	red   = "\033[31m" // ❌ Rot: Fehler oder Warnung
	reset = "\033[0m"  // ⬅️ Farbe zurücksetzen
)

// User repräsentiert die Datenstruktur der Anfrage
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// interner Zähler für Demo-Zwecke (simuliert IDs)
var userIDCounter = 1

// CreateUserHandler verarbeitet POST-Anfragen für /api/v1/users
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Nur POST-Anfragen zulassen
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		log.Printf("%s❌ Ungültige HTTP-Methode:%s %s", red, reset, r.Method)
		return
	}

	// Request-Body in User-Struktur einlesen
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		log.Printf("%s❌ Fehler beim Parsen der Anfrage:%s %v", red, reset, err)
		return
	}

	// Validierung: Email darf nicht leer sein
	if user.Email == "" {
		http.Error(w, "Missing email field", http.StatusBadRequest)
		log.Printf("%s❌ Fehlende E-Mail im Request:%s %+v", red, reset, user)
		return
	}

	// Erfolgreiche Erstellung → Log in grün
	log.Printf("%s✅ Benutzer erstellt:%s Name=%s, Email=%s, ID=%d",
		green, reset, user.Name, user.Email, userIDCounter)

	// JSON-Response vorbereiten
	response := map[string]interface{}{
		"status":  "success",
		"id":      userIDCounter,
		"message": "User created successfully",
	}

	// ID hochzählen (nur Demo, kein persistenter Speicher)
	userIDCounter++

	// Header setzen und JSON zurückgeben
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
