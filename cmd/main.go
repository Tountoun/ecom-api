package main

import (
	"log"

	"github.com/Tountoun/ecom-api/cmd/api"
	"github.com/Tountoun/ecom-api/config"
	_db "github.com/Tountoun/ecom-api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db := _db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUsername,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	
	_db.InitStorage(db)

	server := api.NewAPIServer(":8080", db)

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}