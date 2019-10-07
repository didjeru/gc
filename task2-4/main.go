package main

import (
	"fmt"
)

func main() {
	// Здесь решето возвращает массив из 100 простых чисел
	primes := getFirst100PrimeNumbers(542)

	fmt.Println("Первые", len(primes), "простых чисел:")

	for _, p := range primes {
		fmt.Printf("%v ", p)
	}

	fmt.Println("")
}

func getFirst100PrimeNumbers(n int) []int {
	var primes []int
	arr := make([]bool, n)

	for p := 2; p < n; p++ {
		if arr[p] == true {
			continue
		}

		primes = append(primes, p)

		for k := p * p; k < n; k += p {
			arr[k] = true
		}
	}

	return primes
}
