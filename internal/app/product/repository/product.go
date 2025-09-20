package repository

import (
	"context"
	"fmt"

	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

func (r *RProduct) All(ctx context.Context, keyword string) ([]entity.Product, error) {
	query := `
	SELECT
		id,
		name,
		price,
		created_at,
		updated_at
	FROM
		products 
		`

	var (
		rows *sqlx.Rows
		err  error
	)

	if keyword != "" {
		query += "WHERE name ILIKE $1"
		keyword = fmt.Sprintf("%%%s%%", keyword)
		rows, err = r.db.QueryxContext(ctx, query, keyword)

	} else {
		rows, err = r.db.QueryxContext(ctx, query)
	}

	if rows != nil {
		defer func() {
			if err := rows.Close(); err != nil {
				return
			}
		}()
		// defer rows.Close()
	}

	if err != nil {
		return []entity.Product{}, err
	}

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.StructScan(&product); err != nil {
			return []entity.Product{}, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *RProduct) Get(ctx context.Context, id string) (entity.Product, error) {
	query := `
	SELECT
		id,
		name,
		price,
		created_at,
		updated_at
	FROM
		products
	WHERE
		id = $1
		`

	var product entity.Product

	row := r.db.QueryRowxContext(ctx, query, id)
	if err := row.StructScan(&product); err != nil {
		return entity.Product{}, err
	}

	return product, nil
}
