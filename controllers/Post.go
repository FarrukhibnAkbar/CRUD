package controllers

import(
	"fmt"
	"net/http"
	"io/ioutil"
	db "low/database"
	"encoding/json"
)

func PostProduct(w http.ResponseWriter, r *http.Request) {

	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))

	var newProducts db.PostProductBody

	json.Unmarshal(body, &newProducts)

	encoder.Encode(db.PostProducts(newProducts))
}