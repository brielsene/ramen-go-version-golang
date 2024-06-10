package main

import (
	"ramen-go/database"
	"ramen-go/routes"
)

func main() {
	database.ConnectWithDB()
	// middleware.ErrorHandlingMiddleware()
	routes.HandleRequests()
}
