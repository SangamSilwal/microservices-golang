package handlers

import (
	"log"
	"net/http"

	"github.com/SangamSilwal/microservices-golang/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// We will pass this function in the above based on the method that the user has requested
func (p *Product) getProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProduct()
	err := lp.TOjson(rw)

	if err != nil {
		http.Error(rw, "Unable to get Json object", http.StatusInternalServerError)
	}
}
