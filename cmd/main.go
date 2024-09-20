package main	

import (
	"log"
	"github.com/Stephen882-pixel/golang_ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080",nil)
	if err := server.Run(); err != nil{
		log.Fatal(err)
	}
}