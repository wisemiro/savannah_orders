package models

import "time"

type Order struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	TotalAmout int64     `json:"total_amount"`
	Quantity   int       `json:"quantity"`
	Products   *Product  `json:"products"`
	Customer   *Customer `json:"customer"`
}
