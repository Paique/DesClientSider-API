package routes

import (
	"dcs-api/data"
	"encoding/json"
	"log"
	"net/http"
)

// GetModKeys returns the list of mod keys
func GetModKeys(w http.ResponseWriter, r *http.Request) {
	dbResp := data.GetKeysList()

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
	dbResp := data.GetContraKeyList()

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
