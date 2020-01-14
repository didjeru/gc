package main

import (
	"fmt"
	"./statistic"
)

func main() {
	avg := statistic.Average([]float64{1, 2})
	fmt.Println("Average:", avg)
	sum := statistic.Sum([]float64{1, 2})
	fmt.Println("Sum:", sum)
}
