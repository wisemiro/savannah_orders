package handlers

import (
	"net/http"
	"savannah/internal/models"

	"github.com/go-chi/render"
)

func (repo *HandlerRepo) CreateProduct() http.HandlerFunc {
	type product struct {
		Name  string `json:"product_name"`
		Price int    `json:"price"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData product
		if err := render.DecodeJSON(r.Body, &reqData); err != nil {
			e := NewBadRequestError(ErrorInvalidJSONBody)
			render.Respond(w, r, e)
			return
		}
		err := repo.store.ProductCreate(r.Context(), models.Product{
			Name:  reqData.Name,
			Price: reqData.Price,
		})
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewStatusCreatedResponse(SuccessMessage, nil))
	}
}
func (repo *HandlerRepo) UpdateProduct() http.HandlerFunc {
	type product struct {
		Name  string `json:"product_name"`
		Price int    `json:"price"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData product
		id := repo.GetInt(w, r, "id")

		if err := render.DecodeJSON(r.Body, &reqData); err != nil {
			e := NewBadRequestError(ErrorInvalidJSONBody)
			render.Respond(w, r, e)
			return
		}

		err := repo.store.ProductUpdate(r.Context(), models.Product{
			ID:    id,
			Name:  reqData.Name,
			Price: reqData.Price,
		})
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, nil))
	}
}
func (repo *HandlerRepo) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := repo.GetInt(w, r, "id")
		err := repo.store.ProductDelete(r.Context(), models.Product{ID: id})
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewDeleteResponse(SuccessMessage, nil))

	}
}
func (repo *HandlerRepo) ListProducts() http.HandlerFunc {
	type response struct {
		Total    int               `json:"total"`
		Products []*models.Product `json:"products"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := repo.store.ProductsList(r.Context())
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		data := response{
			Total:    len(products),
			Products: products,
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, data))
	}
}
func (repo *HandlerRepo) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := repo.GetInt(w, r, "id")
		product, err := repo.store.ProductGet(r.Context(), int(id))
		if err != nil {
			e := NewBadRequestError(ErrorProcessingRequest)
			render.Respond(w, r, e)
			return
		}
		render.Respond(w, r, NewStatusOkResponse(SuccessMessage, product))
	}
}
