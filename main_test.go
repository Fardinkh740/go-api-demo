package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Erfolgreicher Test: gültige Daten
func TestCreateUser_Success(t *testing.T) {
	user := User{Name: "Fardin", Email: "fardin@example.com"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	CreateUserHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

// Negativ-Test: Email fehlt
func TestCreateUser_MissingEmail(t *testing.T) {
	user := User{Name: "Fardin"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	CreateUserHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", resp.StatusCode)
	}
}
