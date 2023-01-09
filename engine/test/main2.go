package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct {
}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {

	fmt.Fprintf(w, "Hello there!")
}

func main() {

	mux := http.NewServeMux()

	hello := &helloHandler{}
	mux.Handle("/hello", hello)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
