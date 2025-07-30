package main

import (
	"github.com/spf13/viper"
	"github.com/tomasdepi/koronet-test/core"
)

func main() {

	viper.AutomaticEnv()

	viper.SetDefault("REDIS_HOST", "localhost:6379")
	viper.SetDefault("MYSQL_HOST", "localhost")
	viper.SetDefault("MYSQL_USER", "root")
	viper.SetDefault("MYSQL_PASS", "root")
	viper.SetDefault("MYSQL_DATABASE", "koronet")

	app := core.App{}
	app.Initialize()
	defer app.Finalize()
	app.Run()
}
