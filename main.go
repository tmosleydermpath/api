package main

import (
	_ "expvar"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	httpAddr   = flag.String("http", "0.0.0.0:8080", "HTTP service address.")
	healthAddr = flag.String("health", "0.0.0.0:8090", "HTTP service address.")
)

func main() {
	flag.Parse()

	router := NewRouter()

	log.Println("Starting API Server...")
	log.Printf("HTTP service listening on %s", *httpAddr)
	log.Printf("Health service listening on %s", *healthAddr)

	errChan := make(chan error, 10)

	httpServer := &http.Server{Addr: *httpAddr, Handler: router}
	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", HealthIndex)

	healthServer := &http.Server{Addr: *healthAddr, Handler: Logger(hmux, "Health")}
	go func() {
		errChan <- healthServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			os.Exit(0)
		}
	}

}
