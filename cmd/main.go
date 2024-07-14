package main

import (
	"database/sql"
	"log"

	"github.com/dede182/revesion/cmd/api"
	"github.com/dede182/revesion/config"
	"github.com/dede182/revesion/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMysqlStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		DBName:               config.Envs.DBName,
		Addr:                 config.Envs.DBAddr,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.StartNewServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal("failed to start api server")
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Db connected successfully")
}
