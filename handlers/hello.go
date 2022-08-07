package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello {
	return &hello{l}
}
func (h *hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Printf("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "OOps", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hellow %s", d)
}
