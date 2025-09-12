package main

import (
	"fmt"
	"net/http"

	"demo/4-order-api/configs"
	auth "demo/4-order-api/internal/verify"
)

func main() {
	conf := configs.LoadConfig(".env")

	router := http.NewServeMux()

	auth.NewVerifyHandler(router, auth.VerifyHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server  is listening on port 8081")
	server.ListenAndServe()
}
