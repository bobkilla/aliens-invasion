package planet

type World struct {
	cities map[string]*City
}

func newWorld() *World {
	return &World{cities: make(map[string]*City)}
}

func (world World) GetCity(cityName string) (city *City, found bool) {
	if world.cities[cityName] == nil {
		return nil, false
	}
	return world.cities[cityName], true
}

func (world World) add(city *City) {
	world.cities[city.name] = city
}

func (world World) remove(cityName string) {
	delete(world.cities, cityName)
}

func (world World) CityNames() []string {
	keys := make([]string, 0, len(world.cities))
	for k := range world.cities {
		keys = append(keys, k)
	}
	return keys
}
