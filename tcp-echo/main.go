package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	port              string
	version           string
	requestsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "tcp_echo_requests_processed_total",
		Help: "The total number of processed events",
	})
)

func main() {
	flag.StringVar(&port, "port", "9000", "TCP port")
	flag.StringVar(&version, "version", "v1", "Version to use")
	flag.Parse()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go createTCPListener()
	go createPromEndpoint()

	<-quit
}

func createPromEndpoint() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen: %v\n", err)
	}
}

func createTCPListener() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Server is starting...")

	listener, err := net.Listen("tcp", ":"+port)
	defer listener.Close()
	if err != nil {
		logger.Fatalf("failed to create listener, err: %s\n", err)
	}
	logger.Printf("listening on port %s with version: %s\n", listener.Addr(), version)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatalf("failed to accept connection, err: %s", err)
			continue
		}

		go handleTCPConnection(conn, version)
	}
}

func handleTCPConnection(conn net.Conn, version string) {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Handling new connection...")

	defer func() {
		conn.Close()
		logger.Println("Closing connection...")
	}()

	timeoutDuration := 5 * time.Second

	err := conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	if err != nil {
		logger.Fatalf("failed to set read deadline on connection, err: %s", err)
	}

	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		logger.Fatalf("failed to read data, err: %s", err)
		return
	}

	line := fmt.Sprintf("%s %s", version, bytes)
	logger.Printf("response: %s", line)

	if _, err = conn.Write([]byte(line)); err != nil {
		logger.Fatalf("failed to write response, err: %s", err)
	}

	requestsProcessed.Inc()
}
