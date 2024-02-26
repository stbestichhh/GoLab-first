package main

import (
	"encoding/json"
	"net/http"
)

func TimeHandler(response http.ResponseWriter, request *http.Request) {
	currentTime := GetCurrentTime()
	jsonResponse, err := json.Marshal(currentTime)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
