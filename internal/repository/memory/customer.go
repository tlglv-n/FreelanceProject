package memory

import (
	"context"
	"database/sql"
	"exchanger/internal/domain/customer"
	"github.com/google/uuid"
	"sync"
)

type CustomerRepository struct {
	db map[string]customer.Entity
	sync.RWMutex
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		db: make(map[string]customer.Entity),
	}
}

func (r *CustomerRepository) List(ctx context.Context) (dest []customer.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest = make([]customer.Entity, 0, len(r.db))
	for _, data := range r.db {
		dest = append(dest, data)
	}

	return
}

func (r *CustomerRepository) Add(ctx context.Context, data customer.Entity) (dest string, err error) {
	r.Lock()
	defer r.Unlock()

	id := r.generateID()
	data.ID = id
	r.db[id] = data

	return id, nil
}

func (r *CustomerRepository) Get(ctx context.Context, id string) (dest customer.Entity, err error) {
	r.RLock()
	defer r.RUnlock()

	dest, ok := r.db[id]
	if !ok {
		err = sql.ErrNoRows
		return
	}

	return
}

func (r *CustomerRepository) Update(ctx context.Context, id string, data customer.Entity) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		err = sql.ErrNoRows
	}
	r.db[id] = data

	return
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) (err error) {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.db[id]; !ok {
		err = sql.ErrNoRows
	}
	delete(r.db, id)

	return
}

func (r *CustomerRepository) generateID() string {
	return uuid.New().String()
}
