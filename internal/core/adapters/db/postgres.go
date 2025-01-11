package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresConnection struct {
	DB *sqlx.DB
}

func NewDBConnection(driver, conn string) (*PostgresConnection, error) {
	db, err := sqlx.Connect(driver, conn)
	if err != nil {
		return nil, err
	}

	return &PostgresConnection{DB: db}, nil
}

func (c *PostgresConnection) CloseDBConnection() error {
	if err := c.DB.Close(); err != nil {
		return err
	}
	return nil
}
