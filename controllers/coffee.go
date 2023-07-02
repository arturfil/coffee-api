package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/arturfil/coffee-api/helpers"
	"github.com/arturfil/coffee-api/services"
	"github.com/go-chi/chi/v5"
)

var models services.Models
var coffee = models.Coffee

// GET/coffees
func GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	var coffees services.Coffee
	all, err := coffees.GetAllCoffees()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func GetCoffeeById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	coffee, err := coffee.GetCoffeeById(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffee)
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeResp services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeResp)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffeeResp)
	coffeeCreated, err := coffee.CreateCoffee(coffeeResp)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
	}
}

// TODO check the arguements for the update service
func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var coffee services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, coffee)
	coffeeObj, err := coffee.UpdateCoffee(id, coffee)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		helpers.WriteJSON(w, http.StatusOK, coffeeObj)
	}
}

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := coffee.DeleteCoffee(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	helpers.WriteJSON(w, http.StatusOK, "succesfull deletion")
}
