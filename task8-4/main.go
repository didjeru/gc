package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"./server"
)

func main() {
	srv := server.New(&server.Config{
		Port: ":8080",
	})
	err := srv.Start()
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)
	select {
	case <-ch:
	}
}
