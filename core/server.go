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
	"github.com/tomasdepi/koronet-test/cache"
	"github.com/tomasdepi/koronet-test/rds"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
	redis  *redis.Client
}

func (app *App) Initialize() error {
	// connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DB_USER, DB_PASS, DB_NAME)
	var err error
	app.DB, err = rds.InitDB() //sql.Open("mysql", connectionString)

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

// func StartWebServer() {

// 	mux := http.NewServeMux()
// 	mux.Handle("/", http.HandlerFunc(HelloKoronet))
// 	mux.Handle("/redis", http.HandlerFunc(HelloRedis))
// 	// mux.Handle("/redis-check", http.HandlerFunc(CheckRedis))
// 	mux.Handle("/metrics", promhttp.Handler())

// 	port := "8080"
// 	fmt.Printf("Server listening on http://localhost:%s\n", port)
// 	http.ListenAndServe(":"+port, mux)
// }
