package main

import (
	"fmt"
	"httpfromtcp/parser"
	"io"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Print("Error While trying to listen on port 8080")
	}
	fmt.Printf("Listening on Port:8080")

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error While Accpeting the connection")
		}
		fmt.Print("Client Sucessfully Connected", conn.RemoteAddr().String())
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 1024)
	var request []byte
	for {
		n, err := conn.Read(data)
		if n > 0 {
			request = append(request, data[:n]...)
			// check if we've reached end of headers
			if strings.Contains(string(request), "\r\n\r\n") {
				break
			}
		}
		if err != nil {
			if err == io.EOF {
				break // client closed connection
			}
			fmt.Println("Error reading data:", err)
			return
		}
	}
	parser.ParseRequestLine(string(request))
	parser.ParseHeaders(string(request))
}
