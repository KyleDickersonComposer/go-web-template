package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	Index(0).Render(r.Context(), w)
}

func counterHandler(w http.ResponseWriter, r *http.Request) {
	parseErr := r.ParseForm()
	if parseErr != nil {
		http.Error(w, "couldn't parse the counter form", http.StatusBadRequest)
		return
	}

	formVal := r.FormValue("count")
	count, err := strconv.Atoi(formVal)
	if err != nil {
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	count++

	templ.RenderFragments(r.Context(), w, Index(count), "counter")
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", indexHandler)
	r.Get("/counter", counterHandler)

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
