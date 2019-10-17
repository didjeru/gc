package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Смотри README.md
func main() {
	bs, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(string(bs))
}
