package main

import (
	"fmt"
	"utils"
)

func main() {
	utils.LoadStringfile("C:/Users/michaelssss/GoglandProjects/SmsLisenerServer/src/utils/HHH")
	fmt.Println(utils.Config)
	//http.HandleFunc("/logsms", functions.Logsms)
	//http.HandleFunc("/printsms", functions.Print)
	//http.ListenAndServe(":8080", nil)
}
