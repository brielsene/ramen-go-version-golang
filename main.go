package main

import (
	"ramen-go/database"
	"ramen-go/routes"
)

func main() {
	database.ConnectWithDB()
	routes.HandleRequests()
}
