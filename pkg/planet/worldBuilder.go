package planet

import (
	"fmt"
	"strings"
)

func BuildWorld(lines []string) *World {
	world := newWorld()

	for _, line := range lines {
		cityName, neighbours := parseLine(strings.TrimSpace(line))
		city := getOrCreateCity(world, cityName)

		for i := 0; i < len(neighbours); i++ {
			fillNeighbour(world, city, neighbours[i])
		}
	}
	return world
}

func parseLine(line string) (cityName string, neighbours []string) {
	splitLine := strings.Split(line, " ")
	cityName = splitLine[0]
	neighbours = splitLine[1:]
	return
}

func fillNeighbour(world *World, city *City, neighbour string) {
	splitLine := strings.Split(neighbour, "=")
	if len(splitLine) != 2 {
		//we ignore if there is more than one occurence of the char '='
		fmt.Printf("Error while parsing neighbour, multiple occurence of the char '=', ignoring\n")
		return
	}

	rawDirection := splitLine[0]
	toCityName := splitLine[1]

	direction, err := toDirection(rawDirection)
	if err != nil {
		//we ignore if the direction is not known
		fmt.Printf("Error while parsing direction, ignoring %s\n", rawDirection)
	} else {
		toCity := getOrCreateCity(world, toCityName)
		city.addNeighbour(direction, toCity)
		inverseDirection, err := direction.inverse()
		if err == nil {
			toCity.addNeighbour(inverseDirection, city)
		}
	}
}

func getOrCreateCity(world *World, cityName string) *City {
	city, found := world.GetCity(cityName)
	if !found {
		city = newCity(cityName)
		world.add(city)
	}
	return city
}
