package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
)

func processReceipt(db *InMemDB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var receipt Receipt

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed\n", http.StatusMethodNotAllowed)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
			http.Error(w, "Invalid request payload\n", http.StatusBadRequest)
			return
		}

		id := uuid.NewString()
		value, err := json.Marshal(receipt)
		if err != nil {
			http.Error(w, "Internal server error\n", http.StatusInternalServerError)
			return
		}

		db.Set(id, string(value))
		if err != nil {
			http.Error(w, "Internal server error\n", http.StatusInternalServerError)
			return
		}

		response := ProcessReceiptResponse{Id: id}
		responseAsJSON, _ := json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(responseAsJSON)+"\n")
	}
}

func getPoints(db *InMemDB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed\n", http.StatusMethodNotAllowed)
			return
		}

		path := strings.Split(r.URL.Path, "/")
		id := path[2]
		value, err := db.Get(id)
		if err == badger.ErrKeyNotFound {
			http.Error(w, "Receipt not found\n", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Failed to retrieve receipt\n", http.StatusInternalServerError)
			return
		}

		var receipt Receipt
		json.Unmarshal([]byte(value), &receipt)
		points := CalculatePoints(receipt)
		response := GetPointsResponse{Points: points}
		responseAsJSON, _ := json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, string(responseAsJSON)+"\n")
	}
}

func setupMux(db *InMemDB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/receipts/process", processReceipt(db))
	mux.HandleFunc("/receipts/{id}/points", getPoints(db))
	return mux
}

func HandleRequests(db *InMemDB) {
	mux := setupMux(db)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
