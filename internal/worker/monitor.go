package worker

import (
	"log"
	"time"

	"assignment2/internal/model"
	"assignment2/internal/store"
)

func StartMonitor(repo *store.Repository[string, *model.Task], stop <-chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				stats := map[model.Status]int{}
				for _, t := range repo.GetAll() {
					stats[t.Status]++
				}
				log.Println("Stats:", stats)
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()
}
