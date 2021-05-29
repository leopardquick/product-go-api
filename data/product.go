package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJson(r io.Reader) error {
	jd := json.NewDecoder(r)
	return jd.Decode(p)
}

func GetProduct() Products {
	return productList
}

type Products []*Product

func AddProduct(p *Product) {
	p.ID = getNextid()
	productList = append(productList, p)
}

func UpdateProduct(p *Product, id int) (*Product, error) {
	prodIndex := findProduct(id)

	if prodIndex == -1 {
		return nil, NotFoudError
	}

	p.ID = id
	productList[prodIndex] = p
	return productList[prodIndex], nil
}

func findProduct(id int) int {
	for index, value := range productList {
		if value.ID == id {
			return index
		}
	}
	return -1
}

func getNextid() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func (p *Products) JsonEncoder(rw io.Writer) error {
	je := json.NewEncoder(rw)
	return je.Encode(p)

}

var NotFoudError error = fmt.Errorf("value note found")

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
