package services

import (
	"savannah/cmd/config"
	"savannah/internal/repository/db"
)

type Store interface {
	UserService
	ProductService
	OrderService
}

type SQLStore struct {
	store db.Store
	conf  config.Conf
}

func NewSQLStore(
	store db.Store,
	conf config.Conf,

) (*SQLStore, error) {
	return &SQLStore{
		store: store,
		conf:  conf,
	}, nil
}
