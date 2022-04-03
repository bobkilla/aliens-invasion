package planet

type City struct {
	name       string
	neighbours map[Direction]*City
}

func (city City) Name() string {
	return city.name
}

func NewCity(cityName string) *City {
	return &City{name: cityName, neighbours: make(map[Direction]*City)}
}

func (city City) addNeighbour(direction Direction, neighbour *City) {
	city.neighbours[direction] = neighbour
}

func (city City) GetNeighbour(direction Direction) *City {
	return city.neighbours[direction]
}

func (city City) NeighbourDirections() []Direction {
	keys := make([]Direction, 0, len(city.neighbours))
	for d := range city.neighbours {
		keys = append(keys, d)
	}
	return keys
}

func (city City) HasNeighbour() bool {
	return len(city.neighbours) > 0
}

func (city City) RemoveCityFromNeighbour() {
	for direction, neighbour := range city.neighbours {
		inverseDirection, err := direction.inverse()
		if err == nil {
			delete(neighbour.neighbours, inverseDirection)
		}
	}
}

func (city City) toString() string {
	result := city.name
	for direction, neighbour := range city.neighbours {
		result = result + " " + direction.toString() + "=" + neighbour.name
	}
	return result
}
