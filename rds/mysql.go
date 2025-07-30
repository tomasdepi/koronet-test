package rds

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() *sql.DB {
	dsn := "root:root@tcp(localhost:3306)/koronet" // Replace with env/config in real apps
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to open MySQL connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping MySQL: %v", err)
	}

	fmt.Println("MySQL connection established.")

	return DB
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Error closing MySQL: %v", err)
	}
}
