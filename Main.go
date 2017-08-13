package main

import (
	"./hello2"
	"net/http"
)

func main() {
	http.HandleFunc("/logsms", hello2.Logsms)
	http.ListenAndServe(":8080", nil)
}
