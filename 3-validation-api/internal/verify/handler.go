// Package verify for verify email
package verify

import (
	"fmt"
	"net/http"

	"demo/3-validation-api/configs"
	"demo/3-validation-api/pkg/res"
)

type VerifyHandler struct {
	*configs.Config
}

type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("Send")
		data := VerifyResponse{
			Email: handler.Email,
		}

		res.JSON(w, data, http.StatusCreated)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
