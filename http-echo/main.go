package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	listenAddr        string
	version           string
	requestsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_echo_requests_processed_total",
		Help: "The total number of processed events",
	})
)

func main() {
	flag.StringVar(&listenAddr, "listen-adr", "8080", "server listen address")
	flag.StringVar(&version, "version", "v1", "version")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Server is starting...")

	router := http.NewServeMux()
	router.Handle("/", handler(version))
	router.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", listenAddr),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Printf("Server is ready to handle requests at :%s \n", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

func handler(version string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		//Increment metric for prometheus
		requestsProcessed.Inc()

		//Write http response
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, version)
	})
}
