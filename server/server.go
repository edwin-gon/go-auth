package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SetupServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/helloWorld", helloWorld)
	return mux
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello World",
	})

	fmt.Printf("Hello World")
}

type TokenHandler struct {
}

type TokenRequest struct {
	token string
}

func (th TokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request TokenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
