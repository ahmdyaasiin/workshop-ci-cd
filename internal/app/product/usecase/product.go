package usecase

import (
	"context"

	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/dto"
	"github.com/rs/zerolog/log"
)

func (u *UProduct) All(ctx context.Context, keyword string) ([]dto.ResponseGetProduct, error) {
	p, err := u.rp.All(ctx, keyword)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[PRODUCT][ALL] failed to fetch products")

		return []dto.ResponseGetProduct{}, err
	}

	products := make([]dto.ResponseGetProduct, 0)
	for _, product := range p {
		products = append(products, product.ParseToDTO())
	}

	return products, nil
}

func (u *UProduct) Get(ctx context.Context, id string) (dto.ResponseGetProduct, error) {
	product, err := u.rp.Get(ctx, id)
	if err != nil {
		log.Error().
			Err(err).
			Msg("[PRODUCT][GET] failed to fetch product by ID")

		return dto.ResponseGetProduct{}, err
	}

	return product.ParseToDTO(), nil
}
