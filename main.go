package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/time", TimeHandler)

	port := ":8795"
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
