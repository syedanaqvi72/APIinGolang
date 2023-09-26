package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employees", GetEmployee).Methods("GET")
	r.HandleFunc("/employees/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/employees", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees/{eid}", UpdateEmployeeById).Methods("PUT")
	r.HandleFunc("/employees/{eid}", DeleteEmployee).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
