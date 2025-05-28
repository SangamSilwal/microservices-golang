package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/SangamSilwal/microservices-golang/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello world")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "OOps err", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World page")
	})

	// http.ListenAndServe(":8000", nil)

	l := log.New(os.Stdout, "Product-api", log.LstdFlags)

	g := log.New(os.Stdout, "Test-Case", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(g)
	/*
		ServeMux is an HTTP request multiplexer.
		 It is used for request routing and dispatching.
		 The request routing is based on URL patterns. Each incoming request's
		URL is matched against a list of registered patterns.
		 A handler for the pattern that most closely fits the URL is called.

	*/

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/goodBye", gb)

	http.ListenAndServe(":8000", sm)
}
