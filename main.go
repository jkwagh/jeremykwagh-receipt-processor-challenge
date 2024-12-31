package main

import "fmt"

func main() {

	router := chi.NewRouter()

	router.User(cors.Handler(cors.Options){
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"}
		ExposedHeaders: []string{"Link"}
		}
		})

}
