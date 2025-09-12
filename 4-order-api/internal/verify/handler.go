// Package verify for verify email
package verify

import (
	"net/http"

	"demo/4-order-api/configs"
	"demo/4-order-api/pkg/res"
)

type VerifyHandler struct {
	*configs.Config
}

type VerifyHandlerDeps struct {
	*configs.Config
}

type EmailRequest struct {
	Email string `json:"email"`
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.JSON(w, "data", http.StatusCreated)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res.JSON(w, "data", http.StatusOK)
	}
}
