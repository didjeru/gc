package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
	}
	duration := time.Since(start)
	fmt.Printf("Operation took: %s\n", duration)
}
