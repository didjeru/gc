package chessboard

import (
	"errors"
	"fmt"
	"./models/horse"
	"./models/point"
	"strconv"
)

// ChessBoard шахматная доска
type ChessBoard struct {
	points []point.Point
	horse  horse.Horse
}

func (cb *ChessBoard) init() {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			cb.points = append(cb.points, point.New(x, y))
		}
	}
}

func (cb *ChessBoard) isPointExists(point *point.Point) bool {
	isAvailablePoint := false

	for _, chessBoardPoint := range cb.points {
		if chessBoardPoint.X == point.X && chessBoardPoint.Y == point.Y {
			isAvailablePoint = true
			break
		}
	}

	return isAvailablePoint
}

// SetHorsePosition Размещение коня на доске
func (cb *ChessBoard) SetHorsePosition(x string, y int) error {
	newHorsePoint := converStringPositionToPoint(x, y)

	if !cb.isPointExists(&newHorsePoint) {
		return errors.New("Извините, данная позиция недоступна в данный момент")
	}

	cb.horse.SetPosition(newHorsePoint)

	return nil
}

// GetAvailableHorsePositions Возвращает перечень доступных координат
func (cb *ChessBoard) GetAvailableHorsePositions() []string {
	availablePositions := make([]string, 0)

	for _, point := range cb.horse.GetAvailablePositions() {
		if cb.isPointExists(&point) {
			availablePositions = append(availablePositions, convertPointToTextPosition(point))
		}
	}

	return availablePositions
}

func convertPointToTextPosition(p point.Point) string {
	char := string(p.X + 97)
	return fmt.Sprintf("%s%v", char, p.Y+1)
}

func converStringPositionToPoint(x string, y int) point.Point {
	chars := []rune(x)
	return point.New(int(chars[0])-97, y-1)
}

// ParsePosition разбивает шахматную координату на букву и число
func ParsePosition(str string) (x string, y int, err error) {
	if len(str) > 2 {
		return "", -1, errors.New("Неправильная координата")
	}

	chars := []rune(str)
	y, err = strconv.Atoi(string(chars[1]))

	return string(chars[0]), y, err
}

// New инициализация шахматная доски
func New() ChessBoard {
	horse := horse.New()

	cb := ChessBoard{
		points: make([]point.Point, 0, 64),
		horse:  horse,
	}

	cb.init()

	return cb
}
