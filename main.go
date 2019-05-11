package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lcostea/cfp-api/pkg/conference"
)

func getOpenCallForPapers(w http.ResponseWriter, r *http.Request) {
	cfp := conference.CallForPapers{}
	respondWithJSON(w, http.StatusOK, cfp)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/callforpapers", getOpenCallForPapers).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
