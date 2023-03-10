package main

import (
	json2 "encoding/json"
	"log"
	"net/http"
	"time"
)

func HandleTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var currentTime = time.Now()
	var resp = make(map[string]string)
	resp["time"] = currentTime.Format("2006-01-02T15:04:05Z07:00")
	var json, err = json2.Marshal(resp)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(json)
	return

}
