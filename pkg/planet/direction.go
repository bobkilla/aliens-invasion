package planet

import "errors"

type Direction int64

/*
	Represents all the direction that we want to handle. It is possible to had other direction if needed

*/
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

func toDirection(rawDirection string) (direction Direction, err error) {
	switch rawDirection {
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
