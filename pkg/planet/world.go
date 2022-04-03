package planet

type World struct {
	cities map[string]*City
}

func NewWorld() *World {
	return &World{cities: make(map[string]*City)}
}

func (world World) GetCity(cityName string) (city *City, found bool) {
	if world.cities[cityName] == nil {
		return nil, false
	}
	return world.cities[cityName], true
}

func (world World) Add(city *City) {
	world.cities[city.name] = city
}

func (world World) Remove(cityName string) {
	delete(world.cities, cityName)
}

func (world World) CityNames() []string {
	keys := make([]string, 0, len(world.cities))
	for k := range world.cities {
		keys = append(keys, k)
	}
	return keys
}

func (world World) IsEmpty() bool {
	return len(world.cities) == 0
}

func (world World) ToString() string {
	result := ""
	for _, city := range world.cities {
		result = result + city.toString() + "\n"
	}
	return result
}
