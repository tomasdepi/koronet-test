package core

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/tomasdepi/koronet-test/cache"
	"github.com/tomasdepi/koronet-test/rds"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	redis  *redis.Client
}

func (app *App) Initialize() error {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v",
		viper.GetString("MYSQL_USER"),
		viper.GetString("MYSQL_PASS"),
		viper.GetString("MYSQL_HOST"),
		viper.GetString("MYSQL_DATABASE"),
	)

	var err error
	app.DB, err = rds.InitDB(connectionString)

	if err != nil {
		log.Fatal("err")
		return err
	}

	app.redis = cache.StartRedis()

	app.Router = mux.NewRouter() //.StrictSlash(true)
	app.handleRoutes()
	fmt.Println("App initialized")
	return nil
}

func (app *App) Run() {
	port := "8080"
	fmt.Printf("Server listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, app.Router)
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/", app.HelloKoronet).Methods("get")
	app.Router.HandleFunc("/redis-health", app.CheckRedisHealth).Methods("get")
	app.Router.HandleFunc("/mysql-health", app.CheckMySQLHealth).Methods("get")
	app.Router.Handle("/metrics", promhttp.Handler()).Methods("get")
}

func (app *App) HelloKoronet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintln(w, "Hi Koronet Team.")
}

func (app *App) CheckRedisHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	result := app.redis.Ping(context.Background())
	fmt.Fprintln(w, result)
}

func (app *App) CheckMySQLHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")

	err := app.DB.Ping()
	if err != nil {
		fmt.Fprintln(w, fmt.Sprintf("MySQL connection failed: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "MySQL is alive")
}
