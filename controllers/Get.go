package controllers

import(
	"net/http"
	"low/database"
	"encoding/json"
)


func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(database.GetProducts())
}