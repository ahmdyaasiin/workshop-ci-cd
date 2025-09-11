package entity

import (
	"time"

	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/dto"
)

type Product struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Price     int64     `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (e Product) ParseToDTO() dto.ResponseGetProduct {
	return dto.ResponseGetProduct{
		ID:    e.ID,
		Name:  e.Name,
		Price: e.Price,
	}
}
