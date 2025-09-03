

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Server listening on port 8080...")

	for {
		// Accept a new connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("New client connected:", conn.RemoteAddr())

		// Handle the connection in a goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Stream the bytes received to stdout
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected:", conn.RemoteAddr())
			} else {
				fmt.Println("Read error:", err)
			}
			return
		}

		// Print raw bytes
		fmt.Printf("Received: %q\n", buf[:n])
		fmt.Print("Closing the connection after reading")
		conn.Close()
	}

	parser.ParseRequest()

}
