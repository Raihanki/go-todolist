package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gotodolist?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
