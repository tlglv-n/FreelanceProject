package memory

import (
	"context"
	"database/sql"
	"exchanger/internal/domain/hire"
	"github.com/google/uuid"
	"sync"
)

type HireRepository struct {
	db map[string]hire.Entity
	sync.RWMutex
}

func NewHireRepository() *HireRepository {
	return &HireRepository{
		db: make(map[string]hire.Entity),
	}
}

func (r *HireRepository) List(ctx context.Context) (dest []hire.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]hire.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *HireRepository) Add(ctx context.Context, data hire.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil
}

func (r *HireRepository) Get(ctx context.Context, id string) (dest hire.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = sql.ErrNoRows
		return
	}

	return
}

func (r *HireRepository) Update(ctx context.Context, id string, data hire.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		err = sql.ErrNoRows
	}
	r.db[id] = data

	return
}

func (r *HireRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		err = sql.ErrNoRows
	}
	delete(r.db, id)

	return
}

func (r *HireRepository) generateID() string {
	return uuid.New().String()
}
