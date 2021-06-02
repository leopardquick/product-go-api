package customhandler

import (
	"net/http"
	"strconv"

	"exaple.com/Product/data"
	"github.com/gorilla/mux"
)

// for put request
//swagger:route PUT /product/{id} updateproduct updateProduct
//update the product from the given id
//response:
// 200 : "ok succefully"
func (p *ProductHandler) PutRequest(rw http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	id, withError := strconv.Atoi(val["id"])

	if withError != nil {
		http.Error(rw, "unable to get id", http.StatusBadRequest)
	}
	prod := r.Context().Value(keyProduct{}).(*data.Product)
	prob, err := data.UpdateProduct(prod, id)
	if err != nil {
		http.Error(rw, "not found", http.StatusInternalServerError)
	}
	p.l.Println(prob)
}
