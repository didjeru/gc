package main

import (
	"errors"
	"fmt"
	"math"
)

// GetDiscriminant Находит дискриминант
func GetDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

// GetX1X2 Находит х1 и х2
func GetX1X2(a float64, b float64, c float64) (x1 float64, x2 float64, err error) {
	d := GetDiscriminant(a, b, c)

	if d == 0 {
		x1 = -b / (2 * a)
		x2 = x1
		return x1, x2, nil
	}

	if d > 0 {
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
		return x1, x2, nil
	}

	return -1, -1, errors.New("Корней нет")
}

func main() {
	var a float64 = 2
	var b float64 = 5
	var c float64 = -3

	x1, x2, err := GetX1X2(a, b, c)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("x1=", x1, "x2=", x2)
	}
}
