package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.String("port", "9000", "TCP port")
	prefix := flag.String("version", "v1", "Version to use")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Server is starting...")

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		logger.Fatalf("failed to create listener, err: %s", err)
		os.Exit(1)
	}
	defer listener.Close()
	logger.Printf("listening on port %s with prefix: %s\n", listener.Addr(), *prefix)

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatalf("failed to accept connection, err: %s", err)
			continue
		}

		go handleConnection(conn, prefix)
	}
}

func handleConnection(conn net.Conn, prefix *string) {
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

	logger.Printf("request: %s", bytes)

	line := fmt.Sprintf("%s %s", *prefix, bytes)
	logger.Printf("response: %s", line)

	_, err = conn.Write([]byte(line))
	if err != nil {
		logger.Fatalf("failed to write response, err: %s", err)
	}
}
