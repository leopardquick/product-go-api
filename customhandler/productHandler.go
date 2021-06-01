package customhandler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"exaple.com/Product/data"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusServiceUnavailable)
}

/*
to handle get request
*/
func (p *ProductHandler) GetRequest(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.JsonEncoder(rw)
	if err != nil {
		http.Error(rw, "no data", http.StatusNotFound)
	}
}

/*
to handle post request
*/

func (p *ProductHandler) PostRequest(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("post methode")
	prod := r.Context().Value(keyProduct{}).(*data.Product)
	data.AddProduct(prod)

}

/*for put methode*/

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

type keyProduct struct{}

func (p *ProductHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.l.Println("pass through middleware")
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "cannot martial data", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})

}
