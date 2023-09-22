package handlers

import (
	"net/http"
	"reflect"
	"savannah/internal/services"
	"testing"
)

func TestHandlerRepo_ListProducts(t *testing.T) {
	type fields struct {
		store services.Store
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &HandlerRepo{
				store: tt.fields.store,
			}
			if got := repo.ListProducts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerRepo.ListProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerRepo_CreateProduct(t *testing.T) {
	type fields struct {
		store services.Store
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &HandlerRepo{
				store: tt.fields.store,
			}
			if got := repo.CreateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerRepo.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerRepo_UpdateProduct(t *testing.T) {
	type fields struct {
		store services.Store
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &HandlerRepo{
				store: tt.fields.store,
			}
			if got := repo.UpdateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerRepo.UpdateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerRepo_DeleteProduct(t *testing.T) {
	type fields struct {
		store services.Store
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &HandlerRepo{
				store: tt.fields.store,
			}
			if got := repo.DeleteProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerRepo.DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlerRepo_GetProduct(t *testing.T) {
	type fields struct {
		store services.Store
	}
	tests := []struct {
		name   string
		fields fields
		want   http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &HandlerRepo{
				store: tt.fields.store,
			}
			if got := repo.GetProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandlerRepo.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
