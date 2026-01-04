package model

import "sync"

type Status string

const (
	Pending    Status = "PENDING"
	InProgress Status = "IN_PROGRESS"
	Done       Status = "DONE"
)

type Task struct {
	mu      sync.Mutex
	ID      string `json:"id"`
	Payload string `json:"payload,omitempty"`
	Status  Status `json:"status"`
}

func (t *Task) SetStatus(s Status) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Status = s
}
