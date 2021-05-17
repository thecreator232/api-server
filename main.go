package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Abhishek is writing his code"}`))
	fmt.Println("Response Out")
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Println("Startign server")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
