package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pkyadav73199/BuildingMicroServicesInGo/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// GET Request
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// POST Request
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT Method")
		//expect the id in the Url
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("I got the id : ", id)
		p.updateProductsId(id, w, r)

	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)
	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
	// w.Write(d)
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal the data", http.StatusBadRequest)
		p.l.Fatal(err)
		return
	}

	data.AddProduct(prod)
	p.l.Printf("Prod : %#v", prod)
}

func (p *Products) updateProductsId(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Update Id method")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal the data", http.StatusBadRequest)
		p.l.Fatal(err)
		return
	}
	err = data.UpdateProductsId(id, prod)
	if err != data.ErrProductNotFound {
		http.Error(w, "Product Not Found", http.StatusBadRequest)
	}

}
