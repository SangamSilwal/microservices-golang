package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/SangamSilwal/microservices-golang/data"
)

// Product is a http.Handler
type Product struct {
	l *log.Logger
}

// New product will create a product handler with given logger
func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// The serverHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	//Handling Request for a list of products
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	} else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	} else if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)

		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Println(g)

		if len(g) != 1 {
			p.l.Println("Invalid URL more than one id")
			http.Error(rw, "Invalid Url", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URL more than 2 capture group")
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid unable to convert it to string", idString)
			http.Error(rw, "Invalid URL", http.StatusBadRequest)
			return
		}
		p.l.Println("Got id", id)
		return
	}

	//Catch all
	//If no method is satisfie return an error
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

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Added Product using POST METHOD")

	//Taking reference of the product form the data file
	prod := &data.Product{}

	err := prod.FromJson(r.Body)

	if err != nil {
		http.Error(rw, "Unable  to marshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}
