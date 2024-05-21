package main

import (
	"goaml-api/config"
	"goaml-api/controllers"
	"goaml-api/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/gorilla/mux"
)

func main() {

	r := gin.Default()
	// Connect to the database
	config.ConnectDB()

	// Initialize the router
	//r := mux.NewRouter()

	// Define the endpoints
	r.POST("/login", controllers.Login)
	//r.HandleFunc("/session/login", controllers.Login).Methods("POST")
	//r.HandleFunc("/session/logout", controllers.Logout).Methods("POST")
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
