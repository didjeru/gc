package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(50 * time.Millisecond)
	var spin time.Duration = 10
	time.Sleep(spin * time.Second)
	//	const n = 45
	//	fibN := fibonacci(n)
	//	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Printf("\rDelay for spinner - %d seconds\n", spin)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

/*func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}*/
