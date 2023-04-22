package database

import (
	"fmt"
)

type Client interface {
}

type ClientImpl struct {
}

func NewClientImpl() *ClientImpl {
	return &ClientImpl{}
}

func (d *ClientImpl) Close() error {
	fmt.Println("Not implemented.")
	return nil
}
