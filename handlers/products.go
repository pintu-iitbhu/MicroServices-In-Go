package handlers

import (
	"log"
	"net/http"

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
