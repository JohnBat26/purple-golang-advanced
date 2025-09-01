package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	NewRandomHandler(router)
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server  is listening on port", server.Addr)
	server.ListenAndServe()
}
