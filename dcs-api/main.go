package main

import (
	"dcs-api/data"
	"dcs-api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("DCS API v0.1")

	router := mux.NewRouter()

	_ = data.GetDbInstance()

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			//Logging for Grafana
			data.StoreLog(r.RequestURI, r.RemoteAddr)
			log.Println("Request URI:", r.RequestURI, r.RemoteAddr)
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/keywords", routes.GetModKeys).Methods("GET")
	router.HandleFunc("/contra", routes.GetContraKeys).Methods("GET")

	log.Println("listening on port " + data.AppPort)

	if err := http.ListenAndServe(":"+data.AppPort, router); err != nil {
		log.Fatalf("error while starting the server: %s", err)
		return
	}
}
