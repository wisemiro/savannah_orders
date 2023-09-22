package services

import (
	"context"
	"fmt"
	"log"
	"savannah/internal/models"
	"savannah/internal/repository/db"
)

type OrderService interface {
	OrderCreate(ctx context.Context, order models.Order) error
	OrderUpdate(ctx context.Context, order models.Order) error
	OrderList(ctx context.Context) ([]*models.Order, error)
	OrderListByCustomer(ctx context.Context, customerID int) ([]*models.Order, error)
	OrderGet(ctx context.Context, orderID int) (*models.Order, error)
}

func (s *SQLStore) OrderCreate(ctx context.Context, order models.Order) error {
	corder, err := s.store.CreateOrder(ctx, db.CreateOrderParams{
		TotalAmount: int32(order.TotalAmout),
		ProductID:   order.Products.ID,
		CustomerID:  order.Customer.ID,
		Quantity:    int32(order.Quantity),
	})

	if err != nil {
		return err
	}

	err = sendSMS(&sms{
		Username:    s.conf.Sms.Name,
		PhoneNumber: corder.PhoneNumber.String,
		ShortCode:   s.conf.Sms.Code,
		APIKey:      s.conf.Sms.Key,
		Message:     fmt.Sprintf("Your order of %s has been placed!", corder.ProductName.String),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *SQLStore) OrderUpdate(ctx context.Context, order models.Order) error {
	if err := s.store.UpdateOrder(ctx, db.UpdateOrderParams{
		TotalAmount: int32(order.TotalAmout),
		ID:          int32(order.ID),
	}); err != nil {
		return err
	}
	return nil
}

func (s *SQLStore) OrderList(ctx context.Context) ([]*models.Order, error) {
	orders, err := s.store.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	ordersList := make([]*models.Order, len(orders))
	for i, v := range orders {
		ordersList[i] = &models.Order{
			ID:         int64(v.ID),
			CreatedAt:  v.OrderedAt.Time,
			UpdatedAt:  v.UpdatedAt.Time,
			DeletedAt:  v.DeletedAt.Time,
			TotalAmout: int64(v.TotalAmount),
			Products: &models.Product{
				ID:    int64(v.ID_2.Int32),
				Name:  v.Name.String,
				Price: int(v.Price.Int32),
			},
			Customer: &models.Customer{
				ID:          int64(v.ID_3.Int32),
				PhoneNumber: v.PhoneNumber.String,
			},
		}
	}
	return ordersList, nil
}

func (s *SQLStore) OrderListByCustomer(ctx context.Context, customerID int) ([]*models.Order, error) {
	orders, err := s.store.ListOrdersByCustomer(ctx, int64(customerID))
	if err != nil {
		return nil, err
	}
	ordersList := make([]*models.Order, len(orders))
	for i, v := range orders {
		ordersList[i] = &models.Order{
			ID:         int64(v.ID),
			CreatedAt:  v.OrderedAt.Time,
			UpdatedAt:  v.UpdatedAt.Time,
			DeletedAt:  v.DeletedAt.Time,
			TotalAmout: int64(v.TotalAmount),
			Products: &models.Product{
				ID:    int64(v.ID_2.Int32),
				Name:  v.Name.String,
				Price: int(v.Price.Int32),
			},
			Customer: &models.Customer{
				ID:          int64(v.ID_3.Int32),
				PhoneNumber: v.PhoneNumber.String,
			},
		}
	}
	return ordersList, nil
}

func (s *SQLStore) OrderGet(ctx context.Context, orderID int) (*models.Order, error) {
	order, err := s.store.GetOrder(ctx, int32(orderID))
	if err != nil {
		return nil, err
	}
	return &models.Order{
		ID:         int64(order.ID),
		CreatedAt:  order.OrderedAt.Time,
		UpdatedAt:  order.UpdatedAt.Time,
		DeletedAt:  order.DeletedAt.Time,
		TotalAmout: int64(order.TotalAmount),
		Products: &models.Product{
			ID:    int64(order.ID_2.Int32),
			Name:  order.Name.String,
			Price: int(order.Price.Int32),
		},
		Customer: &models.Customer{
			ID:          int64(order.ID_3.Int32),
			PhoneNumber: order.PhoneNumber.String,
		},
	}, nil
}
