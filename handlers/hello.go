// Writing test cases in golang
package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Implementing Htt p handler interface
type Hello struct {
	l *log.Logger
	//This is for logging info on the CLI
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world Test case")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops sorry", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
