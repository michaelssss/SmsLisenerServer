package functions

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type dbconnection struct {
	sql.DB
}

var Connection = openConnection("root:liangyuming2@@tcp(127.0.0.1:3305)/sms_log?parseTime=true", 3306)

func openConnection(address string, port int) dbconnection {
	connection, err := sql.Open("mysql", address)
	if err != nil {
		fmt.Println(err)
	}
	return dbconnection{*connection}
}
