package main

import (
	"github.com/tomasdepi/koronet-test/cache"
	"github.com/tomasdepi/koronet-test/core"
	"github.com/tomasdepi/koronet-test/rds"
)

func main() {

	redis := cache.StartRedis()

	defer redis.Close()

	mysqlClient := rds.InitDB()

	defer mysqlClient.Close()

	core.StartWebServer()

}
