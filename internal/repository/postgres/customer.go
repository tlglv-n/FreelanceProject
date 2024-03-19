package postgres

import (
	"context"
	"database/sql"
	"errors"
	"exchanger/internal/domain/customer"
	"exchanger/pkg/market"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) List(ctx context.Context) (dest []customer.Entity, err error) {
	query := `
		SELECT id, full_name, pseudonym
		FROM customers
		ORDER BY id`

	err = r.db.SelectContext(ctx, &dest, query)

	return
}

func (r *CustomerRepository) Add(ctx context.Context, data customer.Entity) (id string, err error) {
	query := `
		INSERT INTO customers (full_name, pseudonym)
		VALUES ($1, $2, $3)
		RETURNING id`

	args := []any{data.FullName, data.Pseudonym}

	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *CustomerRepository) Get(ctx context.Context, id string) (dest customer.Entity, err error) {
	query := `
		SELECT id, full_name, pseudonym
		FROM customers
		WHERE id=$1`

	args := []any{id}

	if err = r.db.GetContext(ctx, &dest, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *CustomerRepository) Update(ctx context.Context, id string, data customer.Entity) (err error) {
	sets, args := r.prepareArgs(data)
	if len(args) > 0 {

		args = append(args, id)
		sets = append(sets, "updated_at=CURRENT_TIMESTAMP")
		query := fmt.Sprintf("UPDATE customers SET %r WHERE id=$%d RETURNING id", strings.Join(sets, ", "), len(args))

		if err = r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				err = market.ErrorNotFound
			}
		}
	}

	return
}

func (r *CustomerRepository) prepareArgs(data customer.Entity) (sets []string, args []any) {
	if data.Pseudonym != nil {
		args = append(args, data.Pseudonym)
		sets = append(sets, fmt.Sprintf("pseudonym=$%d", len(args)))
	}

	if data.FullName != nil {
		args = append(args, data.FullName)
		sets = append(sets, fmt.Sprintf("full_name=$%d", len(args)))
	}

	return
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) (err error) {
	query := `
		DELETE FROM customers
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
