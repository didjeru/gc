package main

import (
	"fmt"
	"strconv"
)

func inputData(msg string) string {
	var data string
	fmt.Print(msg)
	fmt.Scanln(&data)
	if data == "exit" {
		panic(nil)
	}
	return data
}

func enterNumber(msg string) float64 {
	for {
		num, err := strconv.ParseFloat(inputData(msg), 32)
		if err == nil {
			return float64(num)
		}
		fmt.Println("Не удалось распознать число")
	}
}

func getNumDel(amount float64) (int, int) {
	d1 := int(amount)
	d2 := (amount - float64(d1)) * 100
	return d1, int(d2)
}

func main() {
	num := enterNumber("Введите число: ")
	_, d2 := getNumDel(num / 3)
	if d2 == 0 {
		fmt.Println("Число", num, "делится без остатка")
	} else {
		fmt.Println("Число", num, "делится на 3 с остатком", d2)
	}
}
