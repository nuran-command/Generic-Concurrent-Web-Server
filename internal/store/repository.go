package store

import "sync"

type Repository[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewRepository[K comparable, V any]() *Repository[K, V] {
	return &Repository[K, V]{data: make(map[K]V)}
}

func (r *Repository[K, V]) Set(key K, value V) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[key] = value
}

func (r *Repository[K, V]) Get(key K) (V, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	val, ok := r.data[key]
	return val, ok
}

func (r *Repository[K, V]) GetAll() []V {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]V, 0, len(r.data))
	for _, v := range r.data {
		out = append(out, v)
	}
	return out
}
