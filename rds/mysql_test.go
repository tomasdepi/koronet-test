package rds

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func initTestViper() {
	viper.Reset() // reset to not override other tests
	viper.AutomaticEnv()

	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_USER", "root")
	viper.SetDefault("MYSQL_PASS", "root")
	viper.SetDefault("MYSQL_DATABASE", "koronet")

}

func TestDBConnection(t *testing.T) {

	initTestViper()

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v",
		viper.GetString("MYSQL_USER"),
		viper.GetString("MYSQL_PASS"),
		viper.GetString("MYSQL_HOST"),
		viper.GetString("MYSQL_DATABASE"),
	)

	DB, err := InitDB(connectionString)
	if err != nil {
		t.Fatalf("Connection Failed: %v", err)
	}

	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		t.Fatalf("Ping failed: %v", err)
	}
}
