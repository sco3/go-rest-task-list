package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
