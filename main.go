// @APIVersion 0.0.1
// @APITitle Case Management API
// @APIDescription Toolkit for interacting with Cases in 2.0.0
// @Contact tmosley@dermpathlab.com
// @TermsOfServiceUrl http://dermpathlab.com/
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
package main

import (
	_ "expvar"
	"fmt"
	"log"
	"net/http"
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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
