package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

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

func getIntNumberFromStdin() int {
	str := getStringFromStdin()
	amount, err := strconv.Atoi(str)

	if err != nil {
		fmt.Print("Ошибка. Попробуйте снова: ")
		return getIntNumberFromStdin()
	}

	return amount
}

func getDeposit() float64 {
	fmt.Print("Введите сумму в рублях, которую вы кладете на депозит: ")
	return getFloatNumberFromStdin()
}

func getPercent() float64 {
	fmt.Print("Введите размер годовой процентной ставки: ")
	return getFloatNumberFromStdin() / 100
}

func getYears() int {
	fmt.Print("Введите срок размещения депозита в годах: ")
	return getIntNumberFromStdin()
}

func getRublesAndPenny(amount float64) (int, int) {
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
	deposit := getDeposit()
	percent := getPercent()
	years := getYears()
	amount := deposit * math.Pow((1+percent), float64(years))

	{
		rubles, penny := getRublesAndPenny(amount)
		fmt.Println(
			"Общая сумма денег по истечению депозита равна:",
			rubles,
			declOfNum(rubles, []string{"рубль", "рубля", "рублей"}),
			penny,
			declOfNum(penny, []string{"копейка", "копейки", "копеек"}),
		)
	}

	{
		rubles, penny := getRublesAndPenny(amount - deposit)
		fmt.Println(
			"Из которых ваш доход составляет:",
			rubles,
			declOfNum(rubles, []string{"рубль", "рубля", "рублей"}),
			penny,
			declOfNum(penny, []string{"копейка", "копейки", "копеек"}),
		)
	}
}
