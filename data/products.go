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

type Products []*Product

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)

}

func UpdateProductsId(id int, p *Product) error {
	_, idx, er := findProduct(id)
	if er != nil {
		return er
	}

	p.ID = id
	productList[idx] = p

	return nil

}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
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
