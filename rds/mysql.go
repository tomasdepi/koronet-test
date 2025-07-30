package rds

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(connectionString string) (*sql.DB, error) {

	var err error

	for i := 0; i < 10; i++ {
		DB, err = sql.Open("mysql", connectionString)
		if err == nil && DB.Ping() == nil {
			log.Println("MySQL connection established.")
			return DB, nil
		}

		log.Println("Waiting for MySQL connection.....")
		time.Sleep(1 * time.Second)
	}

	return nil, err
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Error closing MySQL: %v", err)
	}
}
