package world

type World struct {
	cities map[string]*City
	//aliens     []*Aliens
}

func newWorld() *World {
	return &World{cities: make(map[string]*City)}
}

func (world World) get(cityName string) (city *City, found bool) {
	if world.cities[cityName] == nil {
		return nil, false
	}
	return world.cities[cityName], true
}

func (world World) add(cityName string, city *City) {
	world.cities[cityName] = city
}

func (world World) remove(cityName string) {
	delete(world.cities, cityName)
}
