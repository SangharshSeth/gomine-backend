package main

import (
	"github.com/sangharshseth/gomine-backend/cmd/api"
	"github.com/sangharshseth/gomine-backend/internal/storage"
	"log"
)

func init() {

}

func main() {

	dynamoDBClient, err := storage.NewClient("ap-south-1", "web-developer")
	if err != nil {
		log.Fatal(err)
	}

	server := api.GetAPIServer(":8080", dynamoDBClient)
	err = server.RunServer()
	if err != nil {
		panic(err)
	}
}
