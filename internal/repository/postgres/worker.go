package postgres

import (
	"context"
	"database/sql"
	"errors"
	"exchanger/internal/domain/worker"
	"exchanger/pkg/market"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type WorkerRepository struct {
	db *sqlx.DB
}

func NewWorkerRepository(db *sqlx.DB) *WorkerRepository {
	return &WorkerRepository{
		db: db,
	}
}

func (r *WorkerRepository) List(ctx context.Context) (dest []worker.Entity, err error) {
	query := `
		SELECT id, full_name, position, description
		FROM workers
		ORDER BY id`

	err = r.db.SelectContext(ctx, &dest, query)

	return
}

func (r *WorkerRepository) Add(ctx context.Context, data worker.Entity) (id string, err error) {
	query := `
		INSERT INTO workers (full_name, position, description)
		VALUES ($1, $2, $3)
		RETURNING id`

	args := []any{data.FullName, data.Position, data.Description}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *WorkerRepository) Get(ctx context.Context, id string) (dest worker.Entity, err error) {
	query := `
		SELECT id, full_name, position, description
		FROM workers
		WHERE id=$1`

	args := []any{id}

	if err = r.db.GetContext(ctx, &dest, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *WorkerRepository) Update(ctx context.Context, id string, data worker.Entity) (err error) {
	sets, args := r.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")
		query := fmt.Sprintf("UPDATE workers SET %r WHERE id=$%d RETURNING id", strings.Join(sets, ", "), len(args))

		if err = r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = market.ErrorNotFound
			}
		}
	}

	return
}

func (r *WorkerRepository) prepareArgs(data worker.Entity) (sets []string, args []any) {
	if data.Description != nil {
		args = append(args, data.Description)
		sets = append(sets, fmt.Sprintf("description=$%d", len(args)))
	}

	if data.FullName != nil {
		args = append(args, data.FullName)
		sets = append(sets, fmt.Sprintf("full_name=$%d", len(args)))
	}

	if data.Position != nil {
		args = append(args, data.FullName)
		sets = append(sets, fmt.Sprintf("position=$%d", len(args)))
	}

	return
}

func (r *WorkerRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
		DELETE FROM workers
		WHERE id=$1
		RETURNING id`

	args := []any{id}

	if err = r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}
