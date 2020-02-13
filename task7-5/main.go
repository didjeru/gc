package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getRandomSpeed() int {
	return rand.Intn(5)
}

func countdown(firstCar string, twoCar string) {
	fmt.Println("")
	fmt.Println("The game is start for " + firstCar + " and " + twoCar + "!")
	fmt.Println("")
	time.Sleep(750 * time.Millisecond)
	fmt.Print("3... ")
	time.Sleep(1 * time.Second)
	fmt.Print("2... ")
	time.Sleep(1 * time.Second)
	fmt.Print("1...  ")
	time.Sleep(1 * time.Second)
	fmt.Print("Race!!")
	fmt.Println("")
	time.Sleep(500 * time.Millisecond)
}

func startGame(firstCar string, twoCar string) {
	countdown(firstCar, twoCar)
}

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	//	startGame("first", "second")
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var cars = []int{
		1,
		2,
	}
	for _, car := range cars {
		// увеличиваем счетчик WaitGroup
		wg.Add(1)
		// стартуем горутину для выборки
		go func(position int) {
			// уменьшаем счетчик при завершении горутины
			defer wg.Done()
			// выбираем данные
			for position < 50 {
				var carSpeed = getRandomSpeed()
				position = position + carSpeed
			}
		}(car)
		fmt.Printf("position %i\n", car)
	}

	// ожидаем, пока все запросы не будут завершены
	wg.Wait()

}
