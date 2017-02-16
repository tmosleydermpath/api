package main

import (
	_ "expvar"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	httpAddr = flag.String("http", "0.0.0.0:8080", "HTTP service address.")
)

func main() {
	flag.Parse()

	router := NewRouter()

	// Diable SSL3.0 support
	//config := &tls.Config{
	//	MinVersion: tls.VersionTLS10,
	//}
	//server := &http.Server{Addr: ":10443", Handler: router, TLSConfig: config}
	//err := server.ListenAndServeTLS("server.crt", "server.key")
	log.Println("Starting API Server...")
	log.Printf("HTTP service listening on %s", *httpAddr)
	httpServer := &http.Server{Addr: *httpAddr, Handler: router}
	go func() {
		log.Fatal(httpServer.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown signal received, exiting...")

}
