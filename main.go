package main

import (
	"encoding/json"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"github.com/lcostea/cfp-api/pkg/conference"
	"io/ioutil"
	"net/http"
	"time"
)

func getOpenCallForPapers(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received a call to /callforpapers ")
	var cfps [2]conference.CallForPapers
	cfps[0] = conference.CallForPapers{
		ConferenceID:   "devseccon",
		ConferenceName: "DevSecCon London 2019",
		URL:            "https://www.devseccon.com/call-for-proposals/",
		Deadline:       time.Date(2019, 07, 14, 0, 0, 0, 0, time.UTC),
		Description:    "Making development and delivery more secure, without compromise.",
		Starts:         time.Date(2019, 11, 14, 0, 0, 0, 0, time.UTC),
	}

	cfps[1] = conference.CallForPapers{
		ConferenceID:   "blndevops",
		ConferenceName: "DevOps Days Berlin",
		URL:            "https://www.papercall.io/2019-berlin",
		Deadline:       time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
		Description:    "DevOpsDays is a worldwide series of technical conferences covering topics of working together, software development, IT infrastructure operations, and the intersection between them",
		Starts:         time.Date(2019, 11, 27, 0, 0, 0, 0, time.UTC),
	}

	respondWithJSON(w, http.StatusOK, cfps)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ready(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received a call to /ready ")
	time.Sleep(5 * time.Second)
}

func healthy(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received a call to /healthy ")
	time.Sleep(1 * time.Second)
}

func main() {
	loggerStd := logger.Init("LoggerExample", true, false, ioutil.Discard)
	r := mux.NewRouter()
	r.HandleFunc("/callforpapers", getOpenCallForPapers)
	r.HandleFunc("/ready", ready)
	r.HandleFunc("/healthy", healthy)
	logger.Info("Starting go demo api on port 3000!")
	if err := http.ListenAndServe(":3000", r); err != nil {
		loggerStd.Fatal(err)
	}
}
