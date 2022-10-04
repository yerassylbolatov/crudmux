package handlers

import "github.com/gorilla/mux"

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarsByMake).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", removeCarById).Methods("DELETE")
	return router
}
