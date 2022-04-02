package world

import "errors"

type Direction int64

const (
	North Direction = iota
	South
	East
	West
)

func (direction Direction) String() (string, error) {
	switch direction {
	case North:
		return "north", nil
	case South:
		return "south", nil
	case East:
		return "east", nil
	case West:
		return "west", nil
	}
	return "", errors.New("Direction is unknown")
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
