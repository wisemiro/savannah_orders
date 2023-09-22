package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

const (
	SuccessMessage string = "success"
)

type SuccessResponse struct {
	TimeStamp time.Time   `json:"time_stamp"`
	Message   string      `json:"message"`
	Status    int         `json:"status"`
	Data      interface{} `json:"data"`
}

func NewStatusOkResponse(message string, data any) *SuccessResponse {
	return &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusOK,
		Data:      data,
	}
}

func NewStatusCreatedResponse(message string, data any) *SuccessResponse {
	return &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusCreated,
		Data:      data,
	}
}

func NewDeleteResponse(message string, data any) *SuccessResponse {
	return &SuccessResponse{
		TimeStamp: time.Now(),
		Message:   message,
		Status:    http.StatusOK,
		Data:      data,
	}
}

func (repo *HandlerRepo) GetInt(w http.ResponseWriter, r *http.Request, key string) int64 {
	id := chi.URLParam(r, key)
	value, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		render.Respond(w, r, NewBadRequestError("unable to parse id"))
		return 0
	}
	if id == "" {
		render.Respond(w, r, NewBadRequestError("id is not provided"))
		return 0
	}
	if value == 0 {
		render.Respond(w, r, NewBadRequestError("id provided is not valid"))
		return 0
	}

	return value
}
