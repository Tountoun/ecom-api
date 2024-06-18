package main

import (
	"log"
	"os"

	"github.com/Tountoun/ecom-api/config"
	_db "github.com/Tountoun/ecom-api/db"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
)

func main() {
	db := _db.NewMySQLStorage(mysqlDriver.Config{
		User:                 config.Envs.DBUsername,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "mysql", driver)

	if err != nil {
		log.Fatalln(err)
	}

	arg := os.Args[len(os.Args) - 1] // get last arg of this method run command

	if arg == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalln(err)
		}
	}

	if arg == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalln(err)
		}
	}
}