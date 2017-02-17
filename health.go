package main

import "net/http"

// HealthIndex Return 404 error when going to host root
func HealthIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
