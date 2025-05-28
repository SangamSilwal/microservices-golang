package data

import (
	"encoding/json"
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

// Creating a Products type containg a array of product
type Products []*Product

func (p *Products) TOjson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProduct() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "latte",
		Description: "Freshly prepared",
		Price:       2.45,
		SKU:         "AbcDE",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Freshly prepared",
		Price:       245,
		SKU:         "AbcDE",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
