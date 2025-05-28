package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	//This is called registering the handler to the server mutex
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

	s := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// http.ListenAndServe(":8000", sm)

	//This below code is for gracefull shutDown
	//GraceFull shutDown ensure that our application completes all the remaining task before shutting down the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received Terminated Gracefull shoutdown", sig)

	//A context is the paryt of the program which tells the computer when to stop executing functions in the background
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
