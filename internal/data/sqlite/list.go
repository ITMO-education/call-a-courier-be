package sqlite

import (
	stderrs "errors"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/itmo-education/delivery-backend/internal/domain"
)

const defaultLimit = 10

func (p *Provider) List(req domain.ListContractRequest) ([]domain.Contract, error) {
	if req.Limit <= 0 || req.Limit > 10 {
		req.Limit = defaultLimit
	}

	r, err := p.db.Query(`
		SELECT 
		    address,
		    owner 
		FROM contracts 
		LIMIT $1 OFFSET $s 
`, req.Limit, req.Offset)
	if err != nil {
		return nil, errors.Wrap(err, "error listing contracts")
	}
	defer func() {
		err = stderrs.Join(err, r.Close())
	}()

	contracts := make([]domain.Contract, 0, req.Limit)

	for r.Next() {
		var c domain.Contract
		err = r.Scan(&c.TonAddress, &c.OwnerAddress)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning db response")
		}
		contracts = append(contracts, c)
	}

	return contracts, nil
}
