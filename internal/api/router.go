package api

import "net/http"

func RegisterRoutes(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateTask(w, r)
		case http.MethodGet:
			h.GetTasks(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/tasks/", h.GetTaskByID)
	mux.HandleFunc("/stats", h.GetStats)

	return mux
}
