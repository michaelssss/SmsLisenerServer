package hello2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	fmt.Println(getMessageFromJson(body))
}
func getMessageFromJson(body []byte) Message {
	var jsonMap interface{}
	json.Unmarshal(body, &jsonMap)
	m := jsonMap.(map[string]interface{})
	var message Message
	ReciviedContent := (m["reciviedContent"]).(string)
	From := (m["from"]).(string)
	ReciviedTime := (m["reciviedTime"]).(string)
	message.ReciviedContent = ReciviedContent
	message.From = From
	message.ReciviedTime = ReciviedTime
	return message
}
