package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"exaple.com/Product/customhandler"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//hh := customhandler.NewHome(l)
	ph := customhandler.NewProduct(l)

	//for creating serve mux for our Api
	sm := mux.NewRouter()

	//for creating subrouer for our api
	getRoute := sm.Methods(http.MethodGet).Subrouter()
	getRoute.HandleFunc("/products", ph.GetRequest)

	putRoute := sm.Methods(http.MethodPut).Subrouter()
	putRoute.Use(ph.MiddlewareProductValidation)
	putRoute.HandleFunc("/product/{id:[0-9]+}", ph.PutRequest)

	postRoute := sm.Methods(http.MethodPost).Subrouter()
	postRoute.Use(ph.MiddlewareProductValidation)
	postRoute.HandleFunc("/product", ph.PostRequest)

	// serve docs to server with redoc middle where
	sh := middleware.Redoc(middleware.RedocOpts{SpecURL: "/swagger.yaml"}, nil)
	getRoute.Handle("/docs", sh)
	getRoute.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		fmt.Println(l, "connected sucefully")
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	sigchan := make(chan os.Signal)

	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	server.Shutdown(tc)
	fmt.Print(time.Now(), " shurt down ", sig)
}
