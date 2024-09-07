package problems

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/sangharshseth/gomine-backend/internal/types"
	"log"
	"net/http"
)

type Handler struct {
	store ProblemStore
}

func NewHandler(store ProblemStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (handler *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /problems", handler.GetProblems)
}
func (handler *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
	// Fetch the data from DynamoDB
	dynamoData, err := handler.store.GetItemFromTable(context.TODO(), "problems")
	if err != nil {
		log.Println("Error fetching data from DynamoDB:", err)
		http.Error(w, "Failed to fetch problems", http.StatusInternalServerError)
		return
	}

	// Unmarshal the DynamoDB items into a slice of Problem structs
	var problems []types.Problem
	err = attributevalue.UnmarshalListOfMaps(dynamoData, &problems)
	if err != nil {
		log.Println("Error unmarshaling data from DynamoDB:", err)
		http.Error(w, "Failed to process problems", http.StatusInternalServerError)
		return
	}

	// Return the problems as a JSON response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(problems)
	if err != nil {
		log.Println("Error encoding JSON response:", err)
		http.Error(w, "Failed to encode problems", http.StatusInternalServerError)
	}
}
