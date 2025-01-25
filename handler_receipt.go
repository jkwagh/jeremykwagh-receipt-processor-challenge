package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func handlerCreateReceipt(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Retailer     string `json:"retailer"`
		PurchaseDate string `json:"purchaseDate"`
		PurchaseTime string `json:"purchaseTime"`
		Items        []Item `json:"items"`
		Total        string `json:"total"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}

	purchaseDate, err := time.Parse("2006-01-02", params.PurchaseDate)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing purchaseDate: %v", err))
	}

	purchaseTime, err := time.Parse("15:04", params.PurchaseTime)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing purchaseTime: %v", err))
	}

	receipt := Receipt{
		ID:           uuid.New(),
		Retailer:     params.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Items:        params.Items,
		Total:        params.Total,
	}

	store.Add(receipt)

	type response struct {
		ID string `json:"id"`
	}

	respondWithJSON(w, 200, response{ID: receipt.ID.String()})
}
