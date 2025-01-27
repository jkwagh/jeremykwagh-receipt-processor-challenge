package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// Function to create new receipt to be used in POST route
func handlerCreateReceipt(w http.ResponseWriter, r *http.Request) {
	//Decode Receipt JSON payload
	type parameters struct {
		Retailer     string `json:"retailer"`
		PurchaseDate string `json:"purchaseDate"`
		PurchaseTime string `json:"purchaseTime"`
		Items        []Item `json:"items"`
		Total        string `json:"total"`
		Points       int    `json:"points"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
	}
	//Parse purchase date into time package format
	purchaseDate, err := time.Parse("2006-01-02", params.PurchaseDate)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing purchaseDate: %v", err))
	}
	//Parse purchase time into time package format
	purchaseTime, err := time.Parse("15:04", params.PurchaseTime)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing purchaseTime: %v", err))
	}

	//Create new receipt using JSON package
	receipt := Receipt{
		ID:           uuid.New(),
		Retailer:     params.Retailer,
		PurchaseDate: purchaseDate,
		PurchaseTime: purchaseTime,
		Items:        params.Items,
		Total:        params.Total,
		Points:       32,
	}

	//Calculate receipt points and assign to receipt
	receipt.Points = int32(handlerPoints(receipt))

	//Save new receipt to memory using Add function from store.go fle
	store.Add(receipt)

	//Response and converting to JSON with respondWithJSON
	type response struct {
		ID string `json:"id"`
	}

	respondWithJSON(w, 200, response{ID: receipt.ID.String()})
}

// Function to retrieve receipt based on ID to be used in GET request endpoint
func handlerGetReceipt(w http.ResponseWriter, r *http.Request) {
	//Pull receipt ID from url
	receiptID := chi.URLParam(r, "id")

	//Run GET request and check that receiptID exists
	receipt, exists := store.Get(receiptID)
	if !exists {
		respondWithError(w, 404, "Receipt not found")
		return
	}

	//Pull receipt points for response
	type response struct {
		Points int `json:"points"`
	}
	points := int(receipt.Points)

	//Response converted to JSON with respondWithJSON
	respondWithJSON(w, 201, response{Points: points})
}
