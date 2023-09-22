package services

import (
	"context"
	"savannah/internal/models"
	"savannah/internal/repository/db"
)

type CustomerService interface {
	CustomerCreate(ctx context.Context, customer models.Customer) error
	CustomerUpdate(ctx context.Context, customer models.Customer) error
	CustomerDelete(ctx context.Context, customer models.Customer) error
	CustomerList(ctx context.Context, customer models.Customer) ([]*models.Customer, error)
	CustomerGet(ctx context.Context, phoneNumber string) (*models.Customer, error)
}

func (s *SQLStore) CustomerCreate(ctx context.Context, customer models.Customer) error {
	if err := s.store.CreateCustomer(ctx, db.CreateCustomerParams{
		PhoneNumber: customer.PhoneNumber,
		Token:       customer.AuthToken,
	}); err != nil {
		return err
	}
	return nil
}
func (s *SQLStore) CustomerUpdate(ctx context.Context, customer models.Customer) error {
	if err := s.store.UpdateCustomer(ctx, db.UpdateCustomerParams{
		ID:          int32(customer.ID),
		PhoneNumber: customer.PhoneNumber,
	}); err != nil {
		return err
	}
	return nil
}
func (s *SQLStore) CustomerDelete(ctx context.Context, customer models.Customer) error {
	if err := s.store.DeleteCustomer(ctx, int32(customer.ID)); err != nil {
		return err
	}
	return nil
}
func (s *SQLStore) CustomerList(ctx context.Context, customer models.Customer) ([]*models.Customer, error) {
	customers, err := s.store.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	customersList := make([]*models.Customer, len(customers))
	for i, v := range customers {
		customersList[i] = &models.Customer{
			ID:          int64(v.ID),
			CreatedAt:   v.CreatedAt.Time,
			UpdatedAt:   v.UpdatedAt.Time,
			DeletedAt:   v.DeletedAt.Time,
			PhoneNumber: v.PhoneNumber,
		}
	}
	return customersList, nil
}
func (s *SQLStore) CustomerGet(ctx context.Context, phoneNumber string) (*models.Customer, error) {
	customer, err := s.store.GetCustomer(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID:          int64(customer.ID),
		CreatedAt:   customer.CreatedAt.Time,
		UpdatedAt:   customer.UpdatedAt.Time,
		DeletedAt:   customer.DeletedAt.Time,
		PhoneNumber: customer.PhoneNumber,
		AuthToken:   customer.Token,
	}, nil
}
