package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/jfen9/logoServer/service"
	"net"
	"os"
	"strings"
)

func main() {
	service := ":8124"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// run as a goroutine
		go handleClient(conn)
	}
}

func closeConnection(c net.Conn) {
	if err := c.Close(); err != nil {
		fmt.Println(err)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer closeConnection(conn)

	// initiating handshake
	if _, err := conn.Write([]byte("hello\n")); err != nil {
		fmt.Println(err)
	}

	var buf = make([]byte, 512)

	handler := service.NewHandler()
	ended := false
	for !ended {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(n, string(buf[0:]))
		reader := bufio.NewScanner(bytes.NewReader(buf[0:n]))
		for reader.Scan() {
			cmd := strings.TrimSpace(reader.Text())

			response := handler.Handle(cmd)
			fmt.Print(response)
			if _, err2 := conn.Write([]byte(response)); err2 != nil {
				fmt.Println("writing to connection error:", err2)
				return
			}
			if cmd == "quit" { ended = true }
		}
		if err := reader.Err(); err != nil {
			fmt.Println("reading standard input:", err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}