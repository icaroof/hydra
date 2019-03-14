package main

import (
	"Hydra/hlogger"
	shieldbuilder "Hydra/shieldBuilder"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	http.HandleFunc("/", sroot)
	http.HandleFunc("/raiseShield", raiseShield)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprintf(w, "Welcome to the Hydra software system")
	logger.Println("Received an http", r.Method, "request from root url")
}

func raiseShield(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	logger.Println("Received an http", r.Method, "request from raiseShield url")

	builder := shieldbuilder.NewShieldBuilder()

	queryParam := r.URL.Query().Get("shield")

	if queryParam != "" {
		if strings.Contains(queryParam, "L") {
			builder.RaiseLeft()
		}

		if strings.Contains(queryParam, "R") {
			builder.RaiseRight()
		}

		if strings.Contains(queryParam, "F") {
			builder.RaiseFront()
		}

		if strings.Contains(queryParam, "B") {
			builder.RaiseBack()
		}
	}

	shield := builder.Build()

	fmt.Fprintf(w, "Shields raised:\n")
	fmt.Fprintf(w, shield.String())

}
