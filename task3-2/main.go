package main

import (
	"errors"
	"fmt"
	"os"
)

var stringQueue []string

func push(str string) {
	stringQueue = append(stringQueue, str)
}

func pop() (string, error) {
	if len(stringQueue) == 0 {
		return "", errors.New("Queue is empty")
	}

	str := stringQueue[0]
	stringQueue = stringQueue[1:]

	return str, nil
}

func main() {
	push("First element")
	push("Second element")
	push("Third element")

	for {
		if str, err := pop(); err == nil {
			fmt.Println(str)
		} else {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}
