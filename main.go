package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	router := NewRouter()

	// Diable SSL3.0 support
	//config := &tls.Config{
	//	MinVersion: tls.VersionTLS10,
	//}
	//server := &http.Server{Addr: ":10443", Handler: router, TLSConfig: config}
	//err := server.ListenAndServeTLS("server.crt", "server.key")
	fmt.Println("Listening on port 8080...")
	server := &http.Server{Addr: ":8080", Handler: router}
	go func() {
		log.Fatal(server.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown signal received, exiting...")

}
