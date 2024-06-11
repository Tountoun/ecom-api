package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)


func NewMySQLStorage(cfg mysql.Config) *sql.DB {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func InitStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	log.Println("DB successfully connected")
}