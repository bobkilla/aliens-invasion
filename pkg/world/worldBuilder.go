package world

import (
	"fmt"
	"strings"
)

func buildWorld(lines []string) *World {
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
	if len(splitLine) == 2 {
		rawDirection := splitLine[0]
		toCityName := splitLine[1]

		direction, err := toDirection(rawDirection)
		if err != nil {
			//we ignore if the direction is not known
			fmt.Printf("Error while parsing direction, ignoring %s", rawDirection)
		} else {
			toCity := getOrCreateCity(world, toCityName)
			city.addNeighbour(direction, toCity)
		}
	} else {
		//we ignore if there is more than one occurence of the char '='
		fmt.Printf("Error while parsing neighbour, multiple occurence of the char '=', ignoring")
	}
}

func getOrCreateCity(world *World, cityName string) *City {
	city, found := world.get(cityName)
	if !found {
		city = createNewCity(cityName)
		world.add(cityName, city)
	}
	return city
}
