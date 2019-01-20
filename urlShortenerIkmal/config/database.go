package config

import (
	"database/sql"
	"fmt"
)

func ConnectDB() *sql.DB {
	//var db, err = sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/db_redirects")
	var db, err = sql.Open("mysql", "root:[YOUR_MYSQL_ROOT_PASSWORD]@tcp(127.0.0.1:3306)/db_redirects")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
