package database

import (
	"database/sql"
	"fmt"

	database_query "github.com/mahjong_sharer/pkg/infrastructures/database/query"
)

type Client interface {
	BeginTx() (*sql.Tx, error)
	Close() error

	GetUserQueries() database_query.UserQueryInterface
}

type ClientImpl struct {
	db *sql.DB
}

func NewClientImpl() *ClientImpl {
	return &ClientImpl{}
}

func (d *ClientImpl) BeginTx() (*sql.Tx, error) {
	fmt.Println("Not implemented.")
	return nil, nil
}

func (d *ClientImpl) Close() error {
	fmt.Println("Not implemented.")
	return nil
}

// TODO: if tx is not nil, use tx
func (d *ClientImpl) GetUserQueries() database_query.UserQueryInterface {
	return database_query.New(d.db)
}
