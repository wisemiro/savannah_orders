// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	//
	//
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) error
	CreateOrder(ctx context.Context, arg CreateOrderParams) (*CreateOrderRow, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) error
	//
	//
	CreateUser(ctx context.Context, arg CreateUserParams) error
	//
	//
	DeleteCustomer(ctx context.Context, id int32) error
	//
	//
	DeleteProduct(ctx context.Context, id int32) error
	//
	//
	DeleteUser(ctx context.Context, id int32) error
	GetCustomer(ctx context.Context, phoneNumber string) (*Customers, error)
	//
	//
	GetOrder(ctx context.Context, id int32) (*GetOrderRow, error)
	//
	//
	GetProduct(ctx context.Context, id int32) (*Products, error)
	GetUser(ctx context.Context, userName string) (*Users, error)
	//
	//
	ListCustomers(ctx context.Context) ([]*Customers, error)
	//
	//
	ListOrders(ctx context.Context) ([]*ListOrdersRow, error)
	//
	//
	ListOrdersByCustomer(ctx context.Context, customerID int64) ([]*ListOrdersByCustomerRow, error)
	//
	//
	ListProducts(ctx context.Context) ([]*Products, error)
	//
	//
	ListUsers(ctx context.Context) ([]*Users, error)
	//
	//
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) error
	//
	//
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) error
	//
	//
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error
	//
	//
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
