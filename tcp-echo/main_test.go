package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

func TestTCPConnection(t *testing.T) {
	go main()
	var conn net.Conn
	var err error
	for {
		conn, err = net.Dial("tcp", ":9000")
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	bytesWritten, err := conn.Write([]byte("\n"))
	if err != nil {
		t.Error("could not write payload to server:", err)
	}

	if bytesWritten != 0 {
		if out, _ := bufio.NewReader(conn).ReadString('\n'); err == nil {
			if strings.TrimSpace(out) != "v1" {
				t.Errorf("response did match expected output: got %s want %s", out, "v1")
			}
		}
	}

	err = conn.Close()
	if err != nil {
		t.Error("could not write close connection", err)
	}
}
