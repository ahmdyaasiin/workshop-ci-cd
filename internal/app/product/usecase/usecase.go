package usecase

import "github.com/ahmdyaasiin/workshop-ci-cd/internal/domain/contract"

type UProduct struct {
	rp contract.IRProduct
}

func New(rp contract.IRProduct) contract.IUProduct {
	return &UProduct{
		rp: rp,
	}
}
