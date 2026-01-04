# Generic-Concurrent-Web-Server
Assignment 2 — Generic Concurrent Web Server (Go)

1. Project Overview

This project implements a concurrent HTTP service for managing background tasks using Go generics, goroutines, channels, and mutexes.
It supports:
•	Task submission (POST /tasks)
•	Listing tasks (GET /tasks)
•	Fetching a task by ID (GET /tasks/{id})
•	Server statistics (GET /stats)
•	Worker pool processing
•	Background monitoring
•	Graceful shutdown

2. Project Structure
```
Assignment2/
├── go.mod
├── main.go
├── internal/
│   ├── api/       # HTTP handler and routes
│   ├── model/     # Task model and status
│   ├── queue/     # Generic task queue
│   ├── worker/    # Worker pool and monitor
│   └── store/     # Generic in-memory repository
└── README.md
```

3. Setup & Run
    1.	Make sure you have Go installed (v1.18+).
    2.	Clone or unzip the project folder.
    3.	Open a terminal in the project directory.
    4.	Run the server:
    ```
    go run .
    ```
    Server runs on: http://localhost:8080
4. API Endpoints

1. Create Task
POST /tasks
Content-Type: application/json
Body:
{ "payload": "my first task" }
2. Get All Tasks
   GET /tasks
3. Get Task by ID
   GET /tasks/{id}
4. Get Stats
   GET /stats
5. Features Implemented
   •	Generics: Used for Repository[K,V] and Queue[T]
   •	Concurrency: Worker pool with buffered channel queue
   •	Mutex: Protects shared state for tasks
   •	Monitoring: Logs tasks every 5 seconds
   •	Graceful Shutdown: Stops server, workers, and monitor properly

⸻

6. Testing (Postman / curl)
    1.	Submit tasks via POST /tasks
    2.	View tasks via GET /tasks
    3.	Check individual tasks via GET /tasks/{id}
    4.	View statistics via GET /stats

Example curl commands:
```
# Submit task
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"payload":"my first task"}'

# List all tasks
curl http://localhost:8080/tasks

# Get task by ID
curl http://localhost:8080/tasks/1

# Get stats
curl http://localhost:8080/stats
```