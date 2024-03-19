package postgres

import (
	"context"
	"database/sql"
	"errors"
	"exchanger/internal/domain/hire"
	"exchanger/pkg/market"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type HireRepository struct {
	db *sqlx.DB
}

func NewHireRepository(db *sqlx.DB) *HireRepository {
	return &HireRepository{
		db: db,
	}
}

func (r *HireRepository) List(ctx context.Context) (dest []hire.Entity, err error) {
	query := `
		SELECT id, job_name, amount, description, position, customer_id
		FROM hires
		ORDER BY id`

	err = r.db.SelectContext(ctx, &dest, query)

	return
}

func (r *HireRepository) Add(ctx context.Context, data hire.Entity) (id string, err error) {
	query := `
		INSERT INTO hires (job_name, amount, description, position, customer_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	args := []any{data.JobName, data.Amount, data.Description, data.Position, data.CustomerID}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *HireRepository) Get(ctx context.Context, id string) (dest hire.Entity, err error) {
	query := `
		SELECT id, job_name, amount, description, position, customer_id
		FROM hires
		WHERE id=$1`

	args := []any{id}

	if err = r.db.GetContext(ctx, &dest, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *HireRepository) Update(ctx context.Context, id string, data hire.Entity) (err error) {
	sets, args := r.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")
		query := fmt.Sprintf("UPDATE hires SET %r WHERE id=$%d RETURNING id", strings.Join(sets, ", "), len(args))

		if err = r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = market.ErrorNotFound
			}
		}
	}

	return
}

func (r *HireRepository) prepareArgs(data hire.Entity) (sets []string, args []any) {
	if data.JobName != nil {
		args = append(args, data.JobName)
		sets = append(sets, fmt.Sprintf("job_name=$%d", len(args)))
	}

	if data.Amount != nil {
		args = append(args, data.Amount)
		sets = append(sets, fmt.Sprintf("amount=$%d", len(args)))
	}

	if data.Description != nil {
		args = append(args, data.Description)
		sets = append(sets, fmt.Sprintf("description=$%d", len(args)))
	}

	if data.Position != nil {
		args = append(args, data.Position)
		sets = append(sets, fmt.Sprintf("position=$%d", len(args)))
	}

	return
}

func (r *HireRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
		DELETE FROM hires
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
