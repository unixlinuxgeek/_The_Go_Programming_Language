// Упражнение 8.3

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	tcpAdr, _ := net.ResolveTCPAddr("tcp", "localhost:8001")
	conn, err := net.DialTCP("tcp", nil, tcpAdr)

	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) //
		log.Println("done")
		done <- struct{}{} // сигнал главной go-подпрограмме

	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done
	// Ожидание завершения фоновой go-подпрограммы
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
