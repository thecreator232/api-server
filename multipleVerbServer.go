package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Abhishek is writing his code GET Command"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "Abhishek is writing his code POST Command"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "Abhishek is writing his code "PUT Command"}`))
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Abhishek is writing his code"}`))

	}

	fmt.Println("Response Out")
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Println("Startign server")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
