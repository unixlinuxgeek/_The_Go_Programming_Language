package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	//var tcpConn net.TCPConn
	conn, err := net.Dial("tcp", "localhost:8000")
	//tcpConn = conn.
	//tcpConn.CloseRead()

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
	conn.Close()
	<-done
	// Ожидание завершения фоновой go-подпрограммы
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
