package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks []Task = make([]Task, 0)
var mu sync.Mutex

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("Method: %v\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		fmt.Printf("TBD\n")
		var newTask Task
		json.NewDecoder(r.Body).Decode(&newTask)
		newTask.ID = len(tasks)
		tasks = append(tasks, newTask)
		w.WriteHeader(http.StatusCreated)
	}

}

func main() {
	init := Task{
		ID:    0,
		Title: "Init",
	}
	tasks = append(tasks, init)
	http.HandleFunc("/tasks", HandleTasks)
	http.ListenAndServe(":8080", nil)
}
