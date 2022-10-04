package handlers

import (
	"net/http"
)

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(Vehicles)
}

func returnCarsByMake(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//vars := mux.Vars(r)
	//carM := vars["make"]
	//var cars = &[]Vehicle{}
	//for _, car := range vehicles {
	//	if car.Make == carM {
	//		*cars = append(*cars, car)
	//	}
	//}
	//json.NewEncoder(w).Encode(cars)
}

func removeCarById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//vars := mux.Vars(r)
	//carId, _ := strconv.Atoi(vars["id"])
	//for i, c := range vehicles {
	//	if c.Id == carId {
	//		vehicles = append(vehicles[:i], vehicles[i+1:]...)
	//	}
	//}
	//json.NewEncoder(w).Encode(vehicles)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//var newCar Vehicle
	//json.NewDecoder(r.Body).Decode(&newCar)
	//vehicles = append(vehicles, newCar)
	//json.NewEncoder(w).Encode(vehicles)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//vars := mux.Vars(r)
	//carId, _ := strconv.Atoi(vars["id"])
	//var updateCar Vehicle
	//json.NewDecoder(r.Body).Decode(&updateCar)
	//for i, c := range vehicles {
	//	if c.Id == carId {
	//		vehicles = append(vehicles[:i], vehicles[i+1:]...)
	//		vehicles = append(vehicles, updateCar)
	//	}
	//}
	//json.NewEncoder(w).Encode(vehicles)
}

func returnCarById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//vars := mux.Vars(r)
	//carId, _ := strconv.Atoi(vars["id"])
	//for _, car := range vehicles {
	//	if car.Id == carId {
	//		json.NewEncoder(w).Encode(car)
	//	}
	//}
}
