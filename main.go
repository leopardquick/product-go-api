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
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := customhandler.NewHome(l)
	ph := customhandler.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/product", ph)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
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
