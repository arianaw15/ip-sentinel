package main

import (
	"database/sql"
	"log"

	"github.com/arianaw15/ip-sentinel/cmd/api"
	"github.com/arianaw15/ip-sentinel/config"
	"github.com/arianaw15/ip-sentinel/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewDBConnection(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddress,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Start(); err != nil {
		log.Fatalf("failed to start API server: %v", err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established successfully")
}
