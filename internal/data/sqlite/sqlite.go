package sqlite

import (
	"database/sql"
	"os"
	"path"

	"github.com/godverv/matreshka/resources"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"

	errors "github.com/Red-Sock/trace-errors"
)

const (
	inMemory = "file::memory:?mode=memory&cache=shared"
)

type Provider struct {
	db *sql.DB
}

func New(cfg *resources.Sqlite) (*Provider, error) {
	err := os.MkdirAll(path.Dir(cfg.Path), 0777)
	if err != nil {
		return nil, errors.Wrap(err, "error creating dir for sqlite db")
	}

	conn, err := sql.Open("sqlite", cfg.Path)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database connection")
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "error pinging db")
	}

	err = checkMigrate(conn)
	if err != nil {
		return nil, errors.Wrap(err, "error checking migrations")
	}

	return &Provider{
		db: conn,
	}, nil
}

func checkMigrate(conn *sql.DB) error {
	goose.SetLogger(logrus.StandardLogger())
	err := goose.SetDialect("sqlite")
	if err != nil {
		return errors.Wrap(err, "error setting dialect")
	}

	err = goose.Up(conn, "./migrations")
	if err != nil {
		return errors.Wrap(err, "error performing up")
	}

	return nil
}
