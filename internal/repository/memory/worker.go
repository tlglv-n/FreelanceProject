package memory

import (
	"context"
	"database/sql"
	"exchanger/internal/domain/worker"
	"github.com/google/uuid"
	"sync"
)

type WorkerRepository struct {
	db map[string]worker.Entity
	sync.RWMutex
}

func NewWorkerRepository() *WorkerRepository {
	return &WorkerRepository{
		db: make(map[string]worker.Entity),
	}
}

func (r *WorkerRepository) List(ctx context.Context) (dest []worker.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]worker.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *WorkerRepository) Add(ctx context.Context, data worker.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil
}

func (r *WorkerRepository) Get(ctx context.Context, id string) (dest worker.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = sql.ErrNoRows
		return
	}

	return
}

func (r *WorkerRepository) Update(ctx context.Context, id string, data worker.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		err = sql.ErrNoRows
	}
	r.db[id] = data

	return
}

func (r *WorkerRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		err = sql.ErrNoRows
	}
	delete(r.db, id)

	return
}

func (r *WorkerRepository) generateID() string {
	return uuid.New().String()
}
