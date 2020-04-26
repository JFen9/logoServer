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
			continue
		}
		// run as a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	conn.Write([]byte("hello\n"))

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(buf[0:n]))
	handler := service.NewHandler()
	var reply strings.Builder
	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())
		fmt.Println(cmd)
		if cmd == "quit" {
			break
		}
		response := handler.Handle(cmd)
		reply.WriteString(response)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("reading standard input:", err)
	}
	_, err2 := conn.Write([]byte(reply.String()))
	if err2 != nil {
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}