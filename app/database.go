package app

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp/golangresfullapi")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal("Fail to Ping: ", err.Error())
	} else {
		log.Println("PING")
		log.Println("PING")
		log.Println("PING")
		log.Println("PING")
	}
	return db

}
