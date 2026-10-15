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

type TaskStore struct {
	tasks []Task
	mu    sync.Mutex
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: []Task{
			{ID: 0, Title: "Init"},
		},
	}
}

func (s *TaskStore) HandleTasks(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Printf("Method: %v\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(s.tasks)
	case http.MethodPost:
		fmt.Printf("TBD\n")
		var newTask Task
		json.NewDecoder(r.Body).Decode(&newTask)
		newTask.ID = len(s.tasks)
		s.tasks = append(s.tasks, newTask)
		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	taskStore := NewTaskStore()
	http.HandleFunc("/tasks", taskStore.HandleTasks)
	http.ListenAndServe(":8080", nil)
}
