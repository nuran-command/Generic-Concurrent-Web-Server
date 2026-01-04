package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"

	"assignment2/internal/model"
	"assignment2/internal/queue"
	"assignment2/internal/store"
)

type Handler struct {
	repo      *store.Repository[string, *model.Task]
	queue     *queue.Queue[string]
	submitted atomic.Int64
	completed atomic.Int64
}

func NewHandler(repo *store.Repository[string, *model.Task], q *queue.Queue[string]) *Handler {
	return &Handler{
		repo:  repo,
		queue: q,
	}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body struct {
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	id := strconv.Itoa(int(h.submitted.Add(1)))
	task := &model.Task{
		ID:      id,
		Payload: body.Payload,
		Status:  model.Pending,
	}

	h.repo.Set(id, task)
	h.queue.Push(id)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.repo.GetAll())
}

func (h *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	task, ok := h.repo.Get(id)
	if !ok {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	var inProgress int
	var completed int64

	for _, t := range h.repo.GetAll() {
		switch t.Status {
		case model.InProgress:
			inProgress++
		case model.Done:
			completed++
		}
	}

	json.NewEncoder(w).Encode(map[string]int64{
		"submitted":   h.submitted.Load(),
		"completed":   completed,
		"in_progress": int64(inProgress),
	})
}
