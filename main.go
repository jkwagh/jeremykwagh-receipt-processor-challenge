package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

var store *ReceiptStore

func main() {

	store = NewReceiptStore()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Get("/ready", handlerReadiness)
	router.Get("/err", handlerErr)
	router.Post("/receipts/process", handlerCreateReceipt)
	router.Get("/receipts/{id}/points", handlerGetReceipt)

	srv := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	log.Printf("Server starting on port :8080")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
