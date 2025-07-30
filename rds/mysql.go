package rds

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v",
		viper.GetString("MYSQL_USER"),
		viper.GetString("MYSQL_PASS"),
		viper.GetString("MYSQL_HOST"),
		viper.GetString("MYSQL_NAME"),
	)

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
