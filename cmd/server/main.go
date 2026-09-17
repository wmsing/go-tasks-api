package main

import (
	"log"
	"net/http"
	"os"

	"github.com/portfolio/go-tasks-api/internal/handlers"
	"github.com/portfolio/go-tasks-api/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mem := store.NewMemory()
	tasks := &handlers.Tasks{Store: mem}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", tasks.Health)
	mux.HandleFunc("GET /tasks", tasks.List)
	mux.HandleFunc("POST /tasks", tasks.Create)
	mux.HandleFunc("PATCH /tasks/{id}/complete", tasks.Complete)

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
