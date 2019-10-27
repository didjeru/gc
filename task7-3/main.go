package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func getStringFromStdin() {
	in := bufio.NewScanner(os.Stdin)
	for {
		in.Scan()
		if in.Text() == "exit" {
			os.Exit(1)
		}
	}
}

func main() {
	go getStringFromStdin()
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
