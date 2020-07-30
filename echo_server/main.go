package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err == io.EOF {
		log.Println("Client disconnected")
	}
	if err != nil {
		log.Println("Unexpected error")
	}
	log.Printf("Recived %d bytes: %s\n", len(s), s)

	log.Println("Write data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()

}

func main() {

	listener, err := net.Listen("tcp", ":20080")

	if err != nil {
		log.Fatalln("Unabel to bind to port")
	}

	log.Println("Listening on 0.0.0.0:20080")

	for {
		conn, err := listener.Accept()
		log.Println("Recived connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go echo(conn)
	}
}
