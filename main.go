package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	router := chi.NewRouter()

	// Gives the browsers the functionality of what it can do
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "https://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness) // Connecting the handlerReady to the /ready path
	v1Router.Get("/err", handlerErr)

	router.Mount("/v1", v1Router)
	
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err := srv.ListenAndServe() // Our server stops here and starts listening to http requests
	if err != nil {
		log.Fatal(err) // Display the error
	}
}