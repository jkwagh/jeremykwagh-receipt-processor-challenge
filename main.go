package main

import "fmt"

func main() {

	router := chi.NewRouter()

	router.User(cors.Handler(cors.Options){
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"*"}
		ExposedHeaders: []string{"Link"},
		AllowCredentials:true,
		MaxAge: 300,
		}
		})

	v1Router :=chi.NewRouter()

	V1Router.Get("ready", handlerReadiness)
	v1Router.Get("err", handlerErr)

}
