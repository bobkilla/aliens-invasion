package planet

import (
	"fmt"
	"math/rand"
	"time"
)

var randomGenerator *rand.Rand

func InitInvasion(world *World, maxAliens int) map[int]*Alien {
	nbAliens := maxAliens
	aliens := make(map[int]*Alien, nbAliens)
	cityNames := world.CityNames()
	randomGenerator = initRandomGenerator()

	for i := 0; i < int(maxAliens) && len(cityNames) > 0; i++ {
		city, found := randomizeCity(world, cityNames, int(nbAliens))
		if !found {
			panic("Error initializing aliens")
		}

		alien := &Alien{id: i, city: city}
		if city.hasAlien() {
			destroyCity(world, aliens, city.alien, alien, city)
			//refresh city names since a city was removed
			cityNames = world.CityNames()
		} else {
			fmt.Printf("Setting alien %d to %s\n", alien.id, city.name)
			city.alien = alien
			aliens[i] = alien
		}
	}

	return aliens
}

func Invade(world *World, aliens map[int]*Alien) {
	fmt.Println("Starting invasion")
	keys := getAlienIds(aliens)

	for i := 0; i < 20; i++ {
		for i := 0; i < len(keys); i++ {
			alien := aliens[keys[i]]

			if !alien.city.hasNeighbour() {
				fmt.Printf("Alien %d is trapped\n", alien.id)
				delete(aliens, alien.id)
				keys = getAlienIds(aliens)
				continue
			}

			city := alien.city
			goToCity := randomNeighbour(city)

			fmt.Printf("Moving alien %d to %s\n", alien.id, city.name)
			if goToCity.hasAlien() {
				destroyCity(world, aliens, goToCity.alien, alien, city)
				keys = getAlienIds(aliens)
			} else {
				moveAlienToCity(goToCity, alien)
			}
		}
	}
}

func randomizeCity(world *World, cityNames []string, max int) (*City, bool) {
	var randIndex int
	if max == 0 {
		randIndex = 0
	} else {
		randIndex = randomGenerator.Intn(max)
	}
	cityName := cityNames[randIndex]
	return world.GetCity(cityName)
}

func randomNeighbour(city *City) *City {
	availableDirections := city.neighbourDirections()
	randIndex := randomGenerator.Intn(len(availableDirections))
	randDirection := city.neighbourDirections()[randIndex]

	return city.getNeighbour(randDirection)
}

func getAlienIds(aliens map[int]*Alien) []int {
	keys := make([]int, 0, len(aliens))
	for k := range aliens {
		keys = append(keys, k)
	}
	return keys
}

func moveAlienToCity(city *City, alien *Alien) {
	alien.city.alien = nil
	city.alien = alien
	alien.city = city

}

func initRandomGenerator() *rand.Rand {
	randomSource := rand.NewSource(time.Now().UnixNano())
	return rand.New(randomSource)
}

func destroyCity(world *World, aliens map[int]*Alien, alien1 *Alien, alien2 *Alien, city *City) {
	delete(aliens, alien1.id)
	delete(aliens, alien2.id)
	world.remove(city.name)
	city.removeCityFromNeighbour()
	fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", city.name, alien1.id, alien2.id)
}
