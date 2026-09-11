package main

import (
	"fmt"
	"os"

	"github.com/alazar-09/auth-go-jwt/database/config"
)

func main() {
	db := config.DbConnection()
	config.NewConnection(db)

	fmt.Println("database connected..")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}


}