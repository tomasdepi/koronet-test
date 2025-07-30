package core

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartWebServer() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(HelloKoronet))
	mux.Handle("/redis", http.HandlerFunc(HelloRedis))
	// mux.Handle("/redis-check", http.HandlerFunc(CheckRedis))
	mux.Handle("/metrics", promhttp.Handler())

	port := "8080"
	fmt.Printf("Server listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, mux)
}

func HelloKoronet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintln(w, "Hi Koronet Team.")
}

func HelloRedis(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi Redis.")
}

// func CheckRedis(w http.ResponseWriter, r *http.Request) {
// 	result := cache.redisClint.Ping(ctx)
// 	fmt.Fprintln(w, result)
// }
