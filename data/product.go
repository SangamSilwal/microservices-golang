package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
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
