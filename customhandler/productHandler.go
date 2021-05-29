package customhandler

import (
	"log"
	"net/http"

	"exaple.com/Product/data"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getRequest(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.postRequest(rw, r)
		return
	}

	rw.WriteHeader(http.StatusServiceUnavailable)
}

/*
to handle get request
*/
func (p *ProductHandler) getRequest(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.JsonEncoder(rw)
	if err != nil {
		http.Error(rw, "no data", http.StatusNotFound)
	}
}

/*
to handle post request
*/

func (p *ProductHandler) postRequest(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("post methode")

	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "cannot martial data", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)

}
