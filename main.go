package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"wc-scheduler/actions"
	"wc-scheduler/models"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func scheduleRelay(w http.ResponseWriter, r *http.Request) {

	duration := models.Duration{}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = json.Unmarshal(body, &duration)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// CRON Jobs
	job := actions.NewJob()
	job.TurnOnRelay(duration.From)
	job.TurnOffRelay(duration.To)

	w.WriteHeader(200)
}

func main() {

	router := mux.NewRouter()

	listen := os.Getenv("PORT")

	if listen == "" {
		listen = "9000"
	}

	router.HandleFunc("/schedule/relay", scheduleRelay).Methods("POST", "OPTIONS")

	if err := http.ListenAndServe(":"+listen, gorillah.CombinedLoggingHandler(os.Stdout, router)); err != nil {
		log.Fatal(err)
	}
}
