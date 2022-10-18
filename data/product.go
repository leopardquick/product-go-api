package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"  validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"required,gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
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

func (p *Product) Validation() error {
	validate := validator.New()
	validate.RegisterValidation("sku", skuValidator)
	return validate.Struct(p)
}

func skuValidator(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	result := reg.FindAllString(fl.Field().String(), -1)
	if len(result) != 1 {
		return false
	}
	return true
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
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
