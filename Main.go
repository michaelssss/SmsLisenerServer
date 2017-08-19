package main

import (
	"flag"
	"net/http"

	"functions"
)

func main() {
	var password string
	flag.StringVar(&password, "password", "123456", "password")
	flag.Parse()
	functions.Password = password
	http.HandleFunc("/logsms", functions.Logsms)
	http.HandleFunc("/printsms", functions.Print)
	http.ListenAndServe(":8080", nil)
}
