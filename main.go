package main

import (
	"FinalProject/controllers/authcontroller"
	"FinalProject/controllers/photoscontroller"
	"FinalProject/controllers/users"
	"FinalProject/middlewares"
	"FinalProject/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDatabase()
	routes := mux.NewRouter()

	routes.HandleFunc("/login", authcontroller.Login).Methods("POST")
	routes.HandleFunc("/register", authcontroller.Register).Methods("POST")
	routes.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	
	routes.HandleFunc("/deleteuser", users.DeleteUser).Methods("DELETE")

	photos := routes.PathPrefix("/photos").Subrouter()
	photos.HandleFunc("", photoscontroller.Home).Methods("GET")
	photos.Use(middlewares.Middleware)

	log.Fatal(http.ListenAndServe(":3000", routes))
}