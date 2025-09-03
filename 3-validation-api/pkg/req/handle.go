// Package req for handle body of requests
package req

import (
	"net/http"

	"demo/3-validation-api/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.JSON(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.JSON(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	return &body, nil
}
