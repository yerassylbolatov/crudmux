package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/yerassylbolatov/crudmux/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Db *sql.DB

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	var vehicles []models.Vehicle
	sqlStatement := `SELECT * FROM cars`
	rows, err := Db.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var vehicle models.Vehicle
		if err = rows.Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model, &vehicle.Price); err != nil {
			log.Fatal(err)
		}
		vehicles = append(vehicles, vehicle)
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}

func returnCarsByMake(w http.ResponseWriter, r *http.Request) {
	carMake := mux.Vars(r)["make"]
	sqlStatement := `SELECT * FROM cars WHERE make=$1`
	rows, err := Db.Query(sqlStatement, carMake)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var vehicles []models.Vehicle
	for rows.Next() {
		var vehicle models.Vehicle
		if err = rows.Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model, &vehicle.Price); err != nil {
			log.Fatal(err)
		}
		vehicles = append(vehicles, vehicle)
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}

func removeCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["id"]
	carId, _ := strconv.Atoi(vars)
	sqlStatement := `DELETE FROM cars WHERE id=$1`
	_, err := Db.Exec(sqlStatement, carId)
	if err != nil {
		log.Fatal(err)
	}
	sqlSelectStatement := `SELECT * FROM cars`
	rows, err := Db.Query(sqlSelectStatement)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var vehicles []models.Vehicle
	for rows.Next() {
		var vehicle models.Vehicle
		if err = rows.Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model, &vehicle.Price); err != nil {
			log.Fatal(err)
		}
		vehicles = append(vehicles, vehicle)
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	vehicle := new(models.Vehicle)
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, vehicle); err != nil {
		log.Fatal(err)
	}
	sqlStatement := `INSERT INTO cars (id, make, model, price) VALUES ($1, $2, $3, $4)`
	_, err := Db.Exec(sqlStatement, vehicle.Id, vehicle.Make, vehicle.Model, vehicle.Price)
	if err != nil {
		log.Fatal(err)
	}
	var vehicles []models.Vehicle
	sqlSelectStatement := `SELECT * FROM cars`
	rows, err := Db.Query(sqlSelectStatement)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var car models.Vehicle
		if err = rows.Scan(&car.Id, &car.Make, &car.Model, &car.Price); err != nil {
			log.Fatal(err)
		}
		vehicles = append(vehicles, car)
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	vehicle := new(models.Vehicle)
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, vehicle)

	vars := mux.Vars(r)["id"]
	carId, _ := strconv.Atoi(vars)

	sqlStatement := `UPDATE cars SET id=$1, make=$2, model=$3, price=$4 WHERE id=$5`
	_, err := Db.Exec(sqlStatement, vehicle.Id, vehicle.Make, vehicle.Model, vehicle.Price, carId)
	if err != nil {
		log.Fatal(err)
	}
	sqlSelectStatement := `SELECT * FROM cars`
	rows, err := Db.Query(sqlSelectStatement)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var vehicles []models.Vehicle
	for rows.Next() {
		var vehicle models.Vehicle
		rows.Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model, &vehicle.Price)
		vehicles = append(vehicles, vehicle)
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}

func returnCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, _ := strconv.Atoi(vars["id"])
	var vehicles []models.Vehicle
	sqlStatement := `SELECT * FROM cars WHERE id=$1`
	rows, err := Db.Query(sqlStatement, carId)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var vehicle models.Vehicle
		if err = rows.Scan(&vehicle.Id, &vehicle.Make, &vehicle.Model, &vehicle.Price); err != nil {
			log.Fatal(err)
		}
		vehicles = append(vehicles, vehicle)
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}
