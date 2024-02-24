package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/time", TimeHandler)

port := ":8795"

}
