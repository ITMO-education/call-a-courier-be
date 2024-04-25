package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
	//_ "github.com/mattn/go-sqlite3"

	errors "github.com/Red-Sock/trace-errors"

	"github.com/itmo-education/delivery-backend/internal/data/sqlite/migrations"
)

const (
	inMemory = "file::memory:?mode=memory&cache=shared"
	inFile   = "./data/sqlite-database.db"
)

type Provider struct {
	db *sql.DB
}

func New() (*Provider, error) {
	conn, err := sql.Open("sqlite", inFile)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database connection")
	}

	err = migrations.Migrate(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error performing migrations")
	}

	return &Provider{
		db: conn,
	}, nil
}
