package planet

type City struct {
	name       string
	neighbours map[Direction]*City
	alien      *Alien
}

func newCity(cityName string) *City {
	return &City{name: cityName, neighbours: make(map[Direction]*City)}
}

func (city City) addNeighbour(direction Direction, neighbour *City) {
	city.neighbours[direction] = neighbour
}

func (city City) getNeighbour(direction Direction) *City {
	return city.neighbours[direction]
}

func (city City) neighbourDirections() []Direction {
	keys := make([]Direction, 0, len(city.neighbours))
	for d := range city.neighbours {
		keys = append(keys, d)
	}
	return keys
}

func (city City) hasNeighbour() bool {
	return len(city.neighbours) > 0
}

func (city City) removeCityFromNeighbour() {
	for direction, neighbour := range city.neighbours {
		//fmt.Printf("Removing %s from %s neighbours (direction %s)", neighbour.name,)
		inverseDirection, err := direction.inverse()
		if err == nil {
			delete(neighbour.neighbours, inverseDirection)
		}
	}
}

func (city City) hasAlien() bool {
	return city.alien != nil
}
