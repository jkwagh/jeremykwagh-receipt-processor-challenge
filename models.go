package main

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Receipt struct {
	ID           uuid.UUID `json:"id"`
	Retailer     string    `json:"retailer"`
	PurchaseDate time.Time `json:"purchaseDate"`
	PurchaseTime time.Time `json:"purchaseTime"`
	Items        []Item    `json:"items"`
	Total        string    `json:"total"`
	Points       int32     `json:"points"`
}
