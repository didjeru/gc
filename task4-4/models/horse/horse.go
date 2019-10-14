package horse

import (
	"gc/task4-4/models/point"
)

// Horse Фигура коня
type Horse struct {
	position         point.Point
	availableOffsets []point.Point
}

// SetPosition размещение коня
func (h *Horse) SetPosition(newPosition point.Point) {
	h.position = newPosition
}

// GetAvailablePositions перечень доступных координат относительно текущей
func (h *Horse) GetAvailablePositions() []point.Point {
	availablePositions := make([]point.Point, 0)

	for _, availablePoint := range h.availableOffsets {
		newX := h.position.X + availablePoint.X
		newY := h.position.Y + availablePoint.Y

		if newX >= 0 && newY >= 0 {
			newPosition := point.New(newX, newY)
			availablePositions = append(availablePositions, newPosition)
		}
	}

	return availablePositions
}

// New создание коня
func New() Horse {
	return Horse{
		availableOffsets: []point.Point{
			point.New(-2, -1),
			point.New(-2, 1),
			point.New(-1, -2),
			point.New(-1, 2),
			point.New(2, -1),
			point.New(2, 1),
			point.New(1, -2),
			point.New(1, 2),
		},
	}
}
