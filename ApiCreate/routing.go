package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", GetEmployee).Methods("GET")
	r.HandleFunc("/employees", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employees", CreateEmployee).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
