package main

import (
	"net/http"

	"github.com/edwin-gon/go-auth/server"
)

func main() {
	mux := server.SetupServer()
	http.ListenAndServe(":3000", mux)
}
