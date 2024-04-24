package migrations

import (
	"database/sql"

	errors "github.com/Red-Sock/trace-errors"
)

type migFunc func(tx *sql.DB) error

func Migrate(db *sql.DB) error {
	for _, v := range []migFunc{v1} {
		err := v(db)
		if err != nil {
			return errors.Wrap(err, "error performing migration")
		}
	}

	return nil
}

func v1(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS contracts (
			address VARCHAR(48) PRIMARY KEY, 
			owner VARCHAR(48)
		);
`)
	return err
}
