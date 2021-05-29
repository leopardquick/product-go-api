package customhandler

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		p.l.Println("Put Methode")

		reg := regexp.MustCompile(`[0-9]+`)
		result := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(result) != 1 {
			http.Error(rw, "invalid url", http.StatusBadRequest)
			return
		}

		stringid := result[0][0]

		idint, _ := strconv.Atoi(stringid)

		p.putRequest(idint, rw, r)
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

/*for put methode*/

func (p *ProductHandler) putRequest(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "cannot martial data", http.StatusBadRequest)
		return
	}

	prob, err := data.UpdateProduct(prod, id)

	if err != nil {
		http.Error(rw, "not found", http.StatusNotFound)
	}

	p.l.Println(prob)
}
