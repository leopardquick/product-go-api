package customhandler

import (
	"net/http"

	"exaple.com/Product/data"
)

//swagger:route POST /product postProduct postProduct
//Post the product to the database
func (p *ProductHandler) PostRequest(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("post methode")
	prod := r.Context().Value(keyProduct{}).(*data.Product)
	data.AddProduct(prod)

}
