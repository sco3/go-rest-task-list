package tasks

import (
	"sync"
)

type TaskStore struct {
	tasks []Task
	mu    sync.Mutex
}
