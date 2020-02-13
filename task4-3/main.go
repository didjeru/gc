package main

import (
	"bufio"
	"fmt"
	"./calculator"
	"os"
)

func getStringFromStdin() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	return in.Text()
}

func main() {
	for {
		fmt.Println("Умный калькулятор. Для помощи, наберите help")
		fmt.Print("Введите необходимые вычисления > ")

		input := getStringFromStdin()

		if input == "help" {
			fmt.Println("Помощь по калькулятору:")
			fmt.Println("Функции: sqrt, abs, log, ln, sin, cos, tan, arcsin, arccos, arctan, max, min")
			fmt.Println("Писать без проблеов, например: 5+5 или min(75.5,65.4) или cos(6)")
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Print("Ошибка - ", err, "\n")
		}
	}
}
