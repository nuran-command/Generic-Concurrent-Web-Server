package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"assignment2/internal/api"
	"assignment2/internal/model"
	"assignment2/internal/queue"
	"assignment2/internal/store"
	"assignment2/internal/worker"
)

func main() {
	repo := store.NewRepository[string, *model.Task]()

	q := queue.NewQueue[string](2)

	handler := api.NewHandler(repo, q)
	router := api.RegisterRoutes(handler)

	pool := worker.NewPool(repo)
	pool.Start(2, q.Channel())

	stopMonitor := make(chan struct{})
	worker.StartMonitor(repo, stopMonitor)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		log.Println("Server running on http://localhost:8080")
		log.Printf("Worker pool started with 2 workers, queue size 2")
		server.ListenAndServe()
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	close(stopMonitor)
	q.Close()
	server.Shutdown(context.Background())
}
