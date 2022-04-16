package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/pkg/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	productHandler := handlers.NewProduct(logger)

	mux := http.NewServeMux()
	mux.Handle("/", productHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigalChannel := make(chan os.Signal)
	signal.Notify(sigalChannel, os.Interrupt)
	signal.Notify(sigalChannel, os.Kill)

	sig := <-sigalChannel
	logger.Println("Received terminate, graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
