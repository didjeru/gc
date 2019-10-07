package main

import (
	"fmt"
	"math/big"
)

func fib(n int) *big.Int {
	x, y := big.NewInt(0), big.NewInt(1)

	for i := 0; i <= n-1; i++ {
		x, y = y, x.Add(x, y)
	}

	return x
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, fib(i))
	}
}
