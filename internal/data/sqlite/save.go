package sqlite

import (
	errors "github.com/Red-Sock/trace-errors"

	"github.com/itmo-education/delivery-backend/internal/domain"
)

func (p *Provider) Save(contract domain.Contract) error {
	_, err := p.db.Exec(`
		INSERT INTO contracts 
				(address, owner)
		VALUES  (     $1,   $2)
`, contract.TonAddress, contract.OwnerAddress)
	if err != nil {
		return errors.Wrap(err, "error saving contract")
	}

	return nil
}
