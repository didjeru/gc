package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getStringFromStdin() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func getIntNumberFromStdin() int {
	str := getStringFromStdin()
	amount, err := strconv.Atoi(str)

	if err != nil {
		fmt.Print("Ошибка. Попробуйте снова: ")
		return getIntNumberFromStdin()
	}

	return amount
}

func getNumState(i int) {
	if i%2 == 0 {
		fmt.Println("Число", i, "четное")
	} else {
		fmt.Println("Число", i, "нечетное")
	}
}

func main() {
	fmt.Print("Введите число: ")
	num := getIntNumberFromStdin()
	getNumState(num)
}
