package data

import (
	"github.com/itmo-education/delivery-backend/internal/domain"
)

type Data interface {
	Save(contract domain.Contract) error
	List(contract domain.ListContractRequest) ([]domain.Contract, error)
}
