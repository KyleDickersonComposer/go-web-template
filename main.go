package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	Index().Render(r.Context(), w)
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", indexHandler)

	distDir := http.Dir("./dist")
	fileServer := http.FileServer(distDir)

	r.Handle("/dist/*", http.StripPrefix("/dist/", fileServer))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
