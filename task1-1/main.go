package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const currentDollarToRubleExchangeRate = 65.34

func getAmountInRubles() float64 {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	amount, err := strconv.ParseFloat(in.Text(), 64)

	if err != nil {
		fmt.Print("Вы ввели некорректную сумму, пожалуйста, попробуйте снова: ")
		return getAmountInRubles()
	}

	return amount
}

func getDecAndFrac(amount float64) (int, int) {
	d1 := int(amount)
	d2 := (amount - float64(d1)) * 100
	return d1, int(d2)
}

// https://gist.github.com/chiliec/22e34af2d08a964fc1418908a19b0c15
func declOfNum(number int, titles []string) string {
	cases := []int{2, 0, 1, 1, 1, 2}
	var currentCase int
	if number%100 > 4 && number%100 < 20 {
		currentCase = 2
	} else if number%10 < 5 {
		currentCase = cases[number%10]
	} else {
		currentCase = cases[5]
	}
	return titles[currentCase]
}

func main() {
	rubles, penny := getDecAndFrac(currentDollarToRubleExchangeRate)

	fmt.Println(
		"Текущий курс: 1 доллар =",
		rubles,
		declOfNum(rubles, []string{"рубль", "рубля", "рублей"}),
		penny,
		declOfNum(penny, []string{"копейка", "копейки", "копеек"}),
	)

	fmt.Print("Введите сумму в рублях, которую нужно конвертнуть в доллары: ")

	amount := getAmountInRubles()
	dollars, cents := getDecAndFrac(amount / currentDollarToRubleExchangeRate)

	fmt.Print("Получившиеся сумма: ")

	if dollars > 0 {
		fmt.Print(dollars, declOfNum(dollars, []string{" доллар", " доллара", " долларов"}), " и ")
	}

	fmt.Println(cents, declOfNum(cents, []string{"цент", "цента", "центов"}))
}
