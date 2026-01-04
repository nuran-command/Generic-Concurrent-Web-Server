package worker

import (
	"log"
	"time"

	"assignment2/internal/model"
	"assignment2/internal/store"
)

type Pool struct {
	repo *store.Repository[string, *model.Task]
}

func NewPool(repo *store.Repository[string, *model.Task]) *Pool {
	return &Pool{repo: repo}
}

func (p *Pool) Start(workers int, tasks <-chan string) {
	for i := 0; i < workers; i++ {
		go func() {
			for id := range tasks {
				task, ok := p.repo.Get(id)
				if !ok {
					continue
				}

				log.Printf("Worker processing task %s\n", id)
				task.SetStatus(model.InProgress)

				time.Sleep(2 * time.Second)

				task.SetStatus(model.Done)
				log.Printf("Worker completed task %s\n", id)
			}
		}()
	}
}
