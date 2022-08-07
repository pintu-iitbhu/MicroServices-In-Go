package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type goodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *goodBye {
	return &goodBye{l}
}

func (h *goodBye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Printf("GoodByeee...guy's")
	fmt.Fprintf(w, "Goodbye")
}
