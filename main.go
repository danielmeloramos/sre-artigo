// main.go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests.",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(requestsTotal)
}

func handler(w http.ResponseWriter, r *http.Request) {
	status := "200"
	if r.URL.Path == "/error" {
		status = "500"
	}
	requestsTotal.WithLabelValues(status).Inc()
	fmt.Fprintf(w, "Hello, Prometheus and Grafana!")
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			time.Sleep(5 * time.Second)
			fmt.Println("Processing...")
		}
	}()

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
