package controllers

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	db "low/database"
)

func PutProduct(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))

	var updateProducts db.PutProductBody

	json.Unmarshal(body, &updateProducts)

	encoder.Encode(db.PutProducts(updateProducts))
}