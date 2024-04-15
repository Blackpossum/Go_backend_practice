package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// fmt.Println("hello World!")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("ports incorectlly set in env")
	}
	// make a router using chi package
	router := chi.NewRouter()

	// configure cors for send request in browser
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// hooks up http handler with a new router
	v1Router := chi.NewRouter()
	v1Router.Get("/ready", Handlerr_management)
	v1Router.Get("/error", Handlerr_Error)
	router.Mount("/v1", v1Router)
	// configure router
	srv := &http.Server{Handler: router, Addr: ":" + portString}
	log.Printf("server running on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
