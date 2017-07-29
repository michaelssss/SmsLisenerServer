package main

import (
	"./hello2"
	"net/http"
)

func main() {
	http.HandleFunc("/sayHello", hello2.SayHello)
	http.ListenAndServe(":8080", nil)
}
