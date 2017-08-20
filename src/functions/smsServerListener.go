package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Message struct {
	ReciviedTime    time.Time
	ReciviedContent string
	From            string
}

func Logsms(w http.ResponseWriter, r *http.Request) {
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
	tx, _ := myconnnection.Begin()
	stmt, err := tx.Prepare("INSERT INTO sms_log.sms_logs(`recivied_time`, `recivied_content`, `from`)VALUES (?,?,?);")
	stmt.Exec(message.ReciviedTime, message.ReciviedContent, message.From)
	defer tx.Commit()
	defer stmt.Close()
	if err != nil {
		fmt.Println(err)
	}
	return message
}
