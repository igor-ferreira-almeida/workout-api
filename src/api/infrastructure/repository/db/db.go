package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var
	Connection *sqlx.DB

func init() {
	username := "root"
	password := "123456"
	host := "localhost:3308"
	dbName := "workoutdb"

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, dbName)
	connection, err := sqlx.Connect("mysql", uri)
	if err != nil {
		panic(err)
	}

	connection.SetMaxIdleConns(10)
	Connection = connection
}