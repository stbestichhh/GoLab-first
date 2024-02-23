package main

import "time"

type TimeResponse struct {
	Time string `json:"time"`
}

func GetCurrentTime() TimeResponse {
	currentTime := time.Now().Format(time.RFC3339)
	return TimeResponse{ Time: currentTime }
}
