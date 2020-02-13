package main

import (
	"fmt"
	"log"
	"net/http"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("http://gmail.com")
	}()
	go func() {
		responses <- request("http://ya.ru")
	}()
	go func() {
		responses <- request("http://google.com")
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	_, err := http.Get(hostname)
	if err != nil {
		log.Print(err)
	}
	return hostname
}

func main() {
	fmt.Println(mirroredQuery())
}
