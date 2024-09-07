package api

import (
	"encoding/json"
	"github.com/sangharshseth/gomine-backend/internal/services/problems"
	"github.com/sangharshseth/gomine-backend/internal/storage"
	"log"
	"net/http"

	"github.com/sangharshseth/gomine-backend/internal/middlewares"
)

type ApiServer struct {
	address string
	db      *storage.Client
}

func GetAPIServer(address string, db *storage.Client) *ApiServer {
	return &ApiServer{
		address: address,
		db:      db,
	}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow headers

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
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
		Handler: middlewares.ApiLogger(withCORS(multiplexer)),
	}

	problemsStore := problems.NewStore(s.db)
	problemsHandler := problems.NewHandler(problemsStore)
	problemsHandler.RegisterRoutes(multiplexer)

	log.Printf("Server has started on %s", s.address)

	return server.ListenAndServe()
}
