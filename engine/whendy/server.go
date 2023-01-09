package whendy

import (
	//"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	port string
}

func HandleJsonRequest(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(Products)
	fmt.Fprintf(w, `[1,2,3]`)
}

func (s Server) Start() {
	http.HandleFunc("/json", HandleJsonRequest)

	http.ListenAndServe(":5000", nil)
}

func (s Server) Test() {
	print("test")
}
