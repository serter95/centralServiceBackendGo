package main

import (
	"github.com/gorilla/mux"
)

func routes(router *mux.Router) {
	router.HandleFunc("/search/{criteria}", searchData).Methods("GET")
}
