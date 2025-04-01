package main

import (
	"net/http"
)

func main() {
	mux := SetupServer()
	http.ListenAndServe(":3000", mux)
}
