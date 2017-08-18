package functions

import (
	"net/http"
	"fmt"
	"time"
	"encoding/json"
)

func Print(w http.ResponseWriter, r *http.Request) {
	w.Write(getAllSms())
}
func getAllSms() []byte {
	var myconnnection = Connection.DB
	myconnnection.Begin()
	result, err := myconnnection.Query("select `recivied_time`,`recivied_content`,`from` from sms_log.sms_logs;")
	fmt.Println(result)
	if err != nil {
		fmt.Println(err)
	}
	size := 4
	dataIndex := -1
	messages := make([]Message, size)
	for result.Next() {
		var recivied_time time.Time
		var recivied_content string
		var from string
		result.Scan(&recivied_time, &recivied_content, &from)
		message := Message{recivied_time, recivied_content, from}
		dataIndex = dataIndex + 1
		if (dataIndex >= size) {
			size = size * 2
			tem := make([]Message, size)
			copy(tem, messages)
			messages = tem
		}

		messages[dataIndex] = message
	}
	bytes, err := json.Marshal(messages)
	return bytes
}
