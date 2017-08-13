package hello2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Logsms(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(saveSMS(getMessageFromJson(body)))
}
func getMessageFromJson(body []byte) Message {
	var jsonMap interface{}
	json.Unmarshal(body, &jsonMap)
	m := jsonMap.(map[string]interface{})
	var message Message
	ReciviedContent := (m["reciviedContent"]).(string)
	From := (m["from"]).(string)
	ReciviedTime := (m["reciviedTime"]).(float64)
	message.ReciviedContent = ReciviedContent
	message.From = From
	message.ReciviedTime = time.Unix(0, int64(ReciviedTime)*int64(time.Millisecond))
	return message
}
func saveSMS(message Message) Message {
	var myconnnection = Connection.DB
	myconnnection.Begin()
	result, err := myconnnection.Exec("INSERT INTO sms_log.sms_logs(`recivied_time`, `recivied_content`, `from`)VALUES (?,?,?);", message.ReciviedTime, message.ReciviedContent, message.From)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
	return message
}
