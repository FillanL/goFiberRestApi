package main

import (
	"fmt"
	"goFiberRestApi/routes"
	"goFiberRestApi/server"
	"os"
)

func main(){
	port := os.Getenv("PORT")	// get the port from the environment variable
	if port == "" {
		fmt.Println("PORT is not set")
		port = ":3000"	// set the default port if not specified
	}
	fmt.Println("Starting server on port:", port)
	app := server.SetUpServer()
	routes.SetupRoutes(app)

	
	

	app.Listen(port)
}