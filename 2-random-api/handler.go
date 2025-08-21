package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

type RandomHandler struct{}

func NewRandomHandler(router *http.ServeMux) {
	handler := &RandomHandler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *RandomHandler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		randInt := rand.IntN(6) + 1
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(strconv.Itoa(randInt)))
	}
}
