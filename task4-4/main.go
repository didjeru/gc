package main

import (
	"fmt"
	"./models/chessboard"
	"os"
	"strings"
)

func enterString(msg string) string {
	var data string

	fmt.Print(msg)
	fmt.Scanln(&data)

	if data == "exit" {
		os.Exit(0)
	}

	return data
}

func main() {
	cb := chessboard.New()

	fmt.Println("Задайте первоначальную позицию коня на шахматной доске!")

	for {
		str := enterString("Введите шахматную координату (от a1 до h8): ")
		x, y, err := chessboard.ParsePosition(str)

		if err == nil {
			err = cb.SetHorsePosition(x, y)

			if err == nil {
				renderedPositions := strings.Join(cb.GetAvailableHorsePositions(), ", ")
				fmt.Println("Доступные позиции:", renderedPositions)
				break
			}
		}

		fmt.Println("К сожалению не удалось разместить коня на доске. Попробуйте ещё раз ввести координаты или введите exit для выхода")
	}
}
