package main

import (
	"fmt"
	"go-sample/database"
	"go-sample/routes"
)

func main() {

	database.ConectaComDB()
	fmt.Println("Starting the server...")
	routes.HandleRequest()
}
