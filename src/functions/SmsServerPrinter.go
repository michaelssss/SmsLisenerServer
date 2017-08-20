package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

var Password string = "123456"

func Print(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
	tx, _ := myconnnection.Begin()
	result, err := tx.Query("select `recivied_time`,`recivied_content`,`from` from sms_log.sms_logs;")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer tx.Commit()
	defer result.Close()
	size := 4
	dataIndex := -1
	messages := make([]Message, size)
	for result.Next() {
		var recivied_time time.Time
		var recivied_content string
		var from string
		err1 := result.Scan(&recivied_time, &recivied_content, &from)
		log.Fatal(err1)
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
