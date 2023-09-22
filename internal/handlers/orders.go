package handlers

import (
	"net/http"
	"savannah/internal/models"

	"github.com/go-chi/render"
)

func (repo *HandlerRepo) CreateOrder() http.HandlerFunc {
	type order struct {
		ProductID  int `json:"product_id"`
		CustomerID int `json:"customer_id"`
		Quantity   int `json:"quantity"`
		TotalAmout int `json:"total_amount"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData order
		if err := render.DecodeJSON(r.Body, &reqData); err != nil {
			e := NewBadRequestError(ErrorInvalidJSONBody)
			render.Respond(w, r, e)
			return
		}
		err := repo.store.OrderCreate(r.Context(), models.Order{
			Quantity:   reqData.Quantity,
			TotalAmout: int64(reqData.TotalAmout),
			Products: &models.Product{
				ID: int64(reqData.ProductID),
			},
			Customer: &models.Customer{
				ID: int64(reqData.CustomerID),
			},
		})
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewStatusCreatedResponse(SuccessMessage, nil))
	}
}

func (repo *HandlerRepo) UpdateOrder() http.HandlerFunc {
	type order struct {
		TotalAmout int `json:"total_amount"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData order
		id := repo.GetInt(w, r, "id")

		if err := render.DecodeJSON(r.Body, &reqData); err != nil {
			e := NewBadRequestError(ErrorInvalidJSONBody)
			render.Respond(w, r, e)
			return
		}
		err := repo.store.OrderUpdate(r.Context(), models.Order{
			ID:         id,
			TotalAmout: int64(reqData.TotalAmout),
		})
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, nil))
	}
}

func (repo *HandlerRepo) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := repo.GetInt(w, r, "id")
		order, err := repo.store.OrderGet(r.Context(), int(id))
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, order))
	}
}

func (repo *HandlerRepo) ListOrdersByCustomer() http.HandlerFunc {
	type response struct {
		Total  int             `json:"total"`
		Orders []*models.Order `json:"orders"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		id := repo.GetInt(w, r, "customer_id")

		orders, err := repo.store.OrderListByCustomer(r.Context(), int(id))
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		data := response{
			Total:  len(orders),
			Orders: orders,
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, data))
	}
}

func (repo *HandlerRepo) ListOrders() http.HandlerFunc {
	type response struct {
		Total  int             `json:"total"`
		Orders []*models.Order `json:"orders"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		orders, err := repo.store.OrderList(r.Context())
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		data := response{
			Total:  len(orders),
			Orders: orders,
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, data))
	}
}
