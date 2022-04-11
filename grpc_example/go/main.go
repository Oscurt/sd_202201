package main

import (
	"log"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("400: Bad request at GetItems")
        w.WriteHeader(http.StatusBadRequest)
        return
	}
	name, sucess := r.URL.Query()["name"]
	if !sucess || len(name) < 1 {
		log.Println("204: No content in GetItems")
        w.WriteHeader(http.StatusNoContent)
        return
	}
	// grpc logica
	log.Println("200: OK in GetItems")
    w.WriteHeader(http.StatusOK)
}

func main() {
	GrpcServer();
	http.HandleFunc("/items", GetItems)
	log.Fatal(http.ListenAndServe(":3000", nil))
}