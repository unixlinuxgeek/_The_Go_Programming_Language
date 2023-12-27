// Clock является TCP-сервером, периодически выводящим время.

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		time.Sleep(10000)
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // Например, обрыв соединения
			continue
		}

		handleConn2(conn)
	}

}

func handleConn2(c net.Conn) {
	defer c.Close()

	for {
		time.Sleep(1000)
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
