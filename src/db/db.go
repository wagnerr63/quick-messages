package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDatabase handle the database connection
func ConnectDatabase() *sql.DB {
	connection := "root:root@tcp(127.0.0.1:3306)/quick_messages?parseTime=true"
	db, err := sql.Open("mysql", connection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
