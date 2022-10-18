// Package classification Product API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
//	Schemes: http
//	BasePath: /
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package customhandler

import (
	"context"
	"log"
	"net/http"

	"exaple.com/Product/data"
)

// A list of product returns in the response
//
//swagger:response productResponse
type productResponseWrapper struct {
	//in: body
	Body []data.Product
}

//swagger:parameters updateProduct
type productParamsWrapper struct {
	//id to find the product
	//in : path
	ID int `json:"id"`
}

//swagger:parameters updateProduct postProduct
type productPutRequireWrapper struct {
	//in : body
	//parameter in body required for update
	Body data.Product
}

type ProductHandler struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusServiceUnavailable)
}

// key for context
type keyProduct struct{}

// middleware for get data and validation
func (p *ProductHandler) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p.l.Println("pass through middleware")
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "cannot martial data", http.StatusBadRequest)
			return
		}

		validaionError := prod.Validation()

		if validaionError != nil {
			http.Error(rw, validaionError.Error(), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), keyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})

}
