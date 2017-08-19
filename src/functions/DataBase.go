package functions

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type dbconnection struct {
	sql.DB
}

var Connection dbconnection

func OpenConnection() {
	connection, err := sql.Open("mysql", "root:liangyuming2@@tcp(127.0.0.1:3306)/sms_log?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(5)
	Connection = dbconnection{*connection}
}
