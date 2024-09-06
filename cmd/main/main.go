package main

import (
	"github.com/sangharshseth/gomine-backend/cmd/api"
)

func init() {

}

func main() {
	server := api.GetAPIServer(":8080")
	server.RunServer()
}
