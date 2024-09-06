package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sangharshseth/gomine-backend/internal/middlewares"
)

type ApiServer struct {
	address string
}

func GetAPIServer(address string) *ApiServer {
	return &ApiServer{
		address: address,
	}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (s *ApiServer) RunServer() error {
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("GET /health-check", func(w http.ResponseWriter, r *http.Request) {
		response := &Response{
			Status:  200,
			Message: "OK",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	server := http.Server{
		Addr:    s.address,
		Handler: middlewares.ApiLogger(multiplexer),
	}

	log.Printf("Server has started on %s", s.address)

	return server.ListenAndServe()
}
