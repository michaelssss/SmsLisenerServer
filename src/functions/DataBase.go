package functions

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbconnection struct {
	sql.DB
}

var Connection = OpenConnection()

func OpenConnection() dbconnection {
	connection, err := sql.Open("mysql", "root:liangyuming2@@tcp(127.0.0.1:3306)/sms_log?parseTime=true&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
	}
	connection.SetMaxOpenConns(5)
	connection.Ping()
	return dbconnection{*connection}
}
