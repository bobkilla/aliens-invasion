package planet

import (
	"fmt"
	"strings"
)

func BuildWorld(lines []string) *World {
	world := NewWorld()

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

func fillNeighbour(world *World, fromCity *City, neighbour string) {
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
		return
	}

	toCity := getOrCreateCity(world, toCityName)
	inverseDirection, err := direction.inverse()
	if fromCity.GetNeighbour(direction) != nil || toCity.GetNeighbour(inverseDirection) != nil {
		fmt.Printf("The route between %s and %s is ignored because one (or both) of those cities already have a route using this direction\n", toCityName, fromCity.name)
		return
	}

	fromCity.addNeighbour(direction, toCity)
	if err == nil {
		toCity.addNeighbour(inverseDirection, fromCity)
	}
}

func getOrCreateCity(world *World, cityName string) *City {
	city, found := world.GetCity(cityName)
	if !found {
		city = NewCity(cityName)
		world.Add(city)
	}
	return city
}
