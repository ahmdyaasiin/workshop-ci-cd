package contract

import (
	"context"

	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/dto"
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/entity"
)

type IUProduct interface {
	All(ctx context.Context, keyword string) ([]dto.ResponseGetProduct, error)
	Get(ctx context.Context, id string) (dto.ResponseGetProduct, error)
}

type IRProduct interface {
	All(ctx context.Context, keyword string) ([]entity.Product, error)
	Get(ctx context.Context, id string) (entity.Product, error)
}
