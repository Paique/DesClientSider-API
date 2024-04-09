package main

import (
	"dcs-api/data"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("DCS API v0.1")

	router := mux.NewRouter()

	_ = GetDbInstance()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/keywords", GetModKeys).Methods("GET")
	router.HandleFunc("/contra", GetContraKeys).Methods("GET")

	log.Println("listening on port " + data.AppPort)

	if err := http.ListenAndServe(":"+data.AppPort, router); err != nil {
		log.Fatalf("error while starting the server: %s", err)
		return
	}
}

// GetModKeys returns the list of mod keys
func GetModKeys(w http.ResponseWriter, r *http.Request) {
	dbResp := GetKeysList()

	if dbResp == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(dbResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panicf("error while encoding response: %s", err)
		return
	}
}

// GetContraKeys returns the list of contra keys
func GetContraKeys(w http.ResponseWriter, r *http.Request) {
	dbResp := GetContraKeyList()

	if dbResp == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := json.NewEncoder(w).Encode(dbResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panicf("error while encoding response: %s", err)
		return
	}
}
