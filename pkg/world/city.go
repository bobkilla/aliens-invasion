package world

type City struct {
	name       string
	neighbours map[Direction]*City
	//aliens     []Alien
}

func createNewCity(cityName string) *City {
	return &City{name: cityName, neighbours: make(map[Direction]*City)}
}

func (city City) addNeighbour(direction Direction, neighbour *City) {
	city.neighbours[direction] = neighbour
}

func (city City) getNeighbour(direction Direction) *City {
	return city.neighbours[direction]
}
