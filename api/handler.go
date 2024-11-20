package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/Wassy4/receipt-processor/models"
	"github.com/wassy4/receipt-processor/utils"
)

var receipts = make(map[string]models.Receipt)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	receipts[id] = receipt

	response := models.ProcessReceiptResponse{Id: id}
	w.Header().Set("Content-Type", "application/json")
	json.Marshal(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.Split(r.URL.Path, "/")
	id := path[1]
	receipt := receipts[id]
	if !receipt {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points := utils.CalculatePoints(receipt)
	response := models.GetPointsResponse{Points: points}

	w.Header().Set("Content-Type", "application/json")
	json.Marshal(response)
}

func HandleRequests() {
	http.HandleFunc("/receipts/process", ProcessReceipt)
	http.HandleFunc("/receipts/{id}/points", GetPoints)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
