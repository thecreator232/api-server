package main

import (
	"fmt"
	"log"
	"net/http"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/", ServeHTTP)
	log.Println("Starting server with function")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
