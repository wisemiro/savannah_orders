package handlers

import "savannah/internal/services"

type HandlerRepo struct {
	store services.Store
}

func NewHandlerRepo(
	store services.Store,

) *HandlerRepo {
	return &HandlerRepo{
		store: store,
	}
}
