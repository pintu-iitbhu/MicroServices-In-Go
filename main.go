package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkyadav73199/BuildingMicroServicesInGo/handlers"
)

func main() {

	// env.Parse()

	l := log.New(os.Stdout, "prod-API ", log.LstdFlags)

	// Create the Handlers
	// hh := handlers.NewHello(l)
	// gh := handlers.NewGoodBye(l)

	ph := handlers.NewProducts(l)

	// Create a new Serve Mux and reguster the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProductsId)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)

	// sm.Handle("/products", ph)

	// Create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap the signTerm and Intrupt  and gracefully shutDown
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Printf("Recieved terminate, graceFully ShutDown", sig)

	tx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tx)
}
