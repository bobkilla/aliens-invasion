package planet

import "errors"

type Direction int64

const (
	North Direction = iota
	South
	East
	West
)

func (direction Direction) inverse() (inverseDirection Direction, err error) {
	switch direction {
	case North:
		return South, nil
	case South:
		return North, nil
	case East:
		return West, nil
	case West:
		return East, nil
	}
	err = errors.New("Direction is unknown")
	return
}

func (direction Direction) toString() string {
	switch direction {
	case North:
		return "north"
	case South:
		return "south"
	case East:
		return "east"
	case West:
		return "west"
	}
	return ""
}

func toDirection(str string) (direction Direction, err error) {
	switch str {
	case "north":
		return North, nil
	case "south":
		return South, nil
	case "east":
		return East, nil
	case "west":
		return West, nil
	}
	err = errors.New("Direction is unknown")
	return
}
