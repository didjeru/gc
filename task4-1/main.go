package main

import "fmt"

type vehicle interface {
	move()
}

//структура "Автомобиль"
type car struct{}

//структура "Поезд"
type train struct{}

//структура "Самолет"
type aircraft struct{}

func (c car) move() {
	fmt.Println("Автомобиль шуршит колесами по дороге")
}
func (a train) move() {
	fmt.Println("Локомотив стучит колесами по рельсам")
}
func (a aircraft) move() {
	fmt.Println("Самолет плывет на крыльях в небесах")
}

func main() {
	zil := car{}
	boing := aircraft{}
	vl80 := train{}

	zil.move()
	boing.move()
	vl80.move()
}
