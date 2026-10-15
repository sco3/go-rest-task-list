package main

import (
	"go-rest-task-list/v2/tasks"
	"log"
	"net/http"
)

func main() {
	taskStore := tasks.NewTaskStore()
	http.HandleFunc("/tasks", taskStore.HandleTasks)
	addr := ":8080"
	log.Printf("Start server on %v", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Cannot start server: %v\n", err)
	}
}
