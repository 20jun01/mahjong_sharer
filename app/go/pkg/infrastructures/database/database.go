package database

import (
	"database/sql"
	"fmt"
)

type Client interface {
	BeginTx() (*sql.Tx, error)
	Close() error
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
