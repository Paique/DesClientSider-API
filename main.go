package main

import (
	"dcs-rest-api/data"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	println("DCS API v0.1")

	data.InitVariables()
	CreateDbInstance()

	router := mux.NewRouter()
	router.HandleFunc("/keywords", GetModKeys).Methods("GET")
	router.HandleFunc("/contra", GetContraKeys).Methods("GET")

	println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":"+data.AppPort, router))
}

func GetModKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbResp := GetKeysList()

	if dbResp == nil {
		println("Response from database is null")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(dbResp)
	if err != nil {
		println(err.Error())
	}
}
func GetContraKeys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dbResp := GetContraKeyList()

	if dbResp == nil {
		println("Response from database is null")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(dbResp)

	if err != nil {
		println(err.Error())
	}

}
