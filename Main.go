package main

import (
	"./functions"
	"net/http"
)

func main() {
	http.HandleFunc("/logsms", functions.Logsms)
	http.HandleFunc("/printsms", functions.Print)
	http.ListenAndServe(":8080", nil)
}
