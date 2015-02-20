// @APIVersion 0.0.1
// @APITitle Case Management API
// @APIDescription Toolkit for interacting with Cases in 2.0.0
// @Contact tmosley@dermpathlab.com
// @TermsOfServiceUrl http://dermpathlab.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Println("Listening on 10443...")

	// Diable SSL3.0 support
	config := &tls.Config{MinVersion: tls.VersionTLS10}
	server := &http.Server{Addr: ":10443", Handler: router, TLSConfig: config}
	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
}
