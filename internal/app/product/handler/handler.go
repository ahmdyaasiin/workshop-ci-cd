package handler

import (
	"github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/contract"
	"github.com/gofiber/fiber/v2"
)

type HProduct struct {
	up contract.IUProduct
}

func New(up contract.IUProduct) *HProduct {
	return &HProduct{
		up: up,
	}
}

func (h *HProduct) MountRoutes(router fiber.Router) {
	products := router.Group("/products")

	products.Get("", h.All)
	products.Get(":productID<int>", h.Get)
}
