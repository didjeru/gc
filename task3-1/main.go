package main

import (
	"fmt"
)

type car struct {
	types         string
	model         string
	year          int
	box           int
	isEngineRun   bool
	isWindowsOpen bool
	boxPercent    float32
}

func printCar(car car) {
	fmt.Println("Тип автомобиля:", car.types)
	fmt.Println("Марка автомобиля:", car.model)
	fmt.Println("Год выпуска:", car.year)
	fmt.Println("Объем багажника:", car.box)
	fmt.Println("Двигатель запущен:", car.isEngineRun)
	fmt.Println("Окна открыты:", car.isWindowsOpen)
	fmt.Println("Багажник заполнен на:", car.boxPercent, "%")
}

func main() {
	zilAuto := car{
		types:         "Грузовик",
		model:         "Зил",
		year:          2011,
		box:           500,
		isEngineRun:   true,
		isWindowsOpen: false,
		boxPercent:    50.25,
	}
	gazAuto := car{
		types:         "Легковой",
		model:         "Чайка",
		year:          1988,
		box:           100,
		isEngineRun:   false,
		isWindowsOpen: true,
		boxPercent:    10.08,
	}

	fmt.Println(zilAuto.model)
	zilAuto.model = "Зил-113"
	fmt.Println(zilAuto.model)

	fmt.Println("")
	printCar(zilAuto)
	fmt.Println("")
	printCar(gazAuto)
}
