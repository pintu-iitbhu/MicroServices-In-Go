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

type Products []*Product

func GetProducts() Products {
	return productList
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = []*Product{

	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Forthy milky coffee",
		Price:       2.45,
		SKU:         "ac323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Latte",
		Description: "Forthy milky coffee",
		Price:       2.45,
		SKU:         "ac323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().String(),
	},
}
