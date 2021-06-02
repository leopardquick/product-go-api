package customhandler

import (
	"net/http"

	"exaple.com/Product/data"
)

//swagger:route  GET /products products listProducts
//Return list of products
//responses:
// 200: productResponse
func (p *ProductHandler) GetRequest(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.JsonEncoder(rw)
	if err != nil {
		http.Error(rw, "no data", http.StatusNotFound)
	}
}
