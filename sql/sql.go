package sql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Db() *sql.DB {
	return db
}

func Connect() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASSWD"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "learngo",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("sql.Open:", err)
	}

	pinErr := db.Ping()

	if pinErr != nil {
		log.Fatal("db.Ping:", pinErr)
	}
	fmt.Println("Mysql Connected.")
}
