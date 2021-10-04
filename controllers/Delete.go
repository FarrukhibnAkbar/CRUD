package controllers

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
	db "low/database"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	e := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)

	var dropProduct db.DeleteProductBody

	json.Unmarshal(body, &dropProduct)

	e.Encode(db.DeleteProducts(dropProduct))
}

