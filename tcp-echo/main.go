package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	port := flag.String("port", "9000", "TCP port")
	prefix := flag.String("prefix", "hello", "Prefix to use")
	flag.Parse()

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		fmt.Println("failed to create listener, err:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("listening on port %s with prefix: %s\n", listener.Addr(), *prefix)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}

		go handleConnection(conn, prefix)
	}
}

func handleConnection(conn net.Conn, prefix *string) {
	fmt.Println("Handling new connection...")

	defer func() {
		conn.Close()
		fmt.Println("Closing connection...")
	}()

	timeoutDuration := 5 * time.Second

	err := conn.SetReadDeadline(time.Now().Add(timeoutDuration))
	if err != nil {
		fmt.Println("failed to set read deadline on connection, err:", err)
	}

	bytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		fmt.Println("failed to read data, err:", err)
		return
	}

	fmt.Printf("request: %s", bytes)

	line := fmt.Sprintf("%s %s", *prefix, bytes)
	fmt.Printf("response: %s", line)

	_, err = conn.Write([]byte(line))
	if err != nil {
		fmt.Println("failed to write response, err:", err)
	}
}
