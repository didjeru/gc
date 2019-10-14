package point

// Point координата
type Point struct {
	X int
	Y int
}

// New Создание координаты
func New(x int, y int) Point {
	return Point{x, y}
}
