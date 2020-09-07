package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"hology/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	conn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", conn)
	if err != nil {
		panic("connection error")
	}

	if db.Ping() != nil {
		panic("connection failed")
	}
}

func CreateConn() *sql.DB {
	return db
}
