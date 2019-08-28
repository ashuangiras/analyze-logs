package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//healthCheck : health check for server
func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("HealthCheck invoked")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All is well :)"))
	return
}

// analyzeLogs : Analyze handler to analyze the logs and create the response
// Method : Post
// Body : logs
// Response : json
func analyzeLogs(w http.ResponseWriter, r *http.Request) {
	log.Println("Log Analyzer invoked . . .")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	log.Println(string(body))
	logLines := strings.Split(string(body), "\n")

	var results Analysis
	for _, line := range logLines {
		result := matchPattern(line)
		log.Println(result)

		if (Result{}) != result {
			results.Results = append(results.Results, result)
		}
	}

	log.Println(results)
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	analyzed, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "error in marshalling", http.StatusInternalServerError)

	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(analyzed)
	return
}
