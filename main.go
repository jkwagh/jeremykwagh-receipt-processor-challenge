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

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/receipts/process", handlerCreateReceipt)
	v1Router.Get("/receipts/{id}/points", handlerGetReceipt)

	router.Mount("/v1", v1Router)

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
