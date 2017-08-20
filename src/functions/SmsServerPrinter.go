package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var Password string = "123456"

func Print(w http.ResponseWriter, r *http.Request) {
	pw := r.RequestURI
	pa, err := url.Parse(pw)
	ps, err := url.ParseQuery(pa.RawQuery)
	if nil != err {
		fmt.Println(err)
	}
	if ps.Get("password") == Password {
		w.Write(getAllSms())
	}

}
func getAllSms() []byte {
	var myconnnection = Connection.DB
	result, err := myconnnection.Query("select `recivied_time`,`recivied_content`,`from` from sms_log.sms_logs;")
	defer result.Close()
	if err != nil {
		fmt.Println(err.Error())
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
		if dataIndex >= size {
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
