package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getSquare(a float64, b float64) float64 {
	return (a * b) / 2
}

func getHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func getPerimetr(a float64, b float64, c float64) float64 {
	return a + b + c
}

func getStringFromStdin() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func getFloatNumberFromStdin() float64 {
	str := getStringFromStdin()
	amount, err := strconv.ParseFloat(str, 64)

	if err != nil {
		fmt.Print("Ошибка. Попробуйте снова: ")
		return getFloatNumberFromStdin()
	}

	return amount
}

func main() {
	fmt.Print("Введите длину первого катета: ")
	a := getFloatNumberFromStdin()

	fmt.Print("Введите длину второго катета: ")
	b := getFloatNumberFromStdin()

	c := getHypotenuse(a, b)

	fmt.Println("Характеристики треугольника:")

	fmt.Println("Длина гипотенузы:", c)
	fmt.Println("Площадь:", getSquare(a, b))
	fmt.Println("Периметр:", getPerimetr(a, b, c))
}
