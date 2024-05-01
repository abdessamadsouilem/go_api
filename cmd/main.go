package main

import (
	"database/sql"
	"log"

	"github.com/abdessamadsouilem/go_api/cmd/api"
	"github.com/abdessamadsouilem/go_api/config"
	"github.com/abdessamadsouilem/go_api/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMysalStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		AllowNativePasswords: true,
		Net:                  "tcp",
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DB: successfully connected")

}
