package invasion

import (
	"fmt"
	"math/rand"
	"time"

	"com.invasion/first/pkg/planet"
)

var randomGenerator *rand.Rand
var citiesWithAliens map[string]*Alien

/*
	This simulator will start by an initialisation phase: every alien will be sent to a city.
	At every point of the simulation, if an alien is sent to a city that already contains an alien,
	this city will be destroyed just like the two aliens.

	The second phase is the simulation of the invasion. At every iteration, every living alien will
	move to another city using routes between cities.
	After 10 000 iterations, the simulation is finished. The aliens probably died of old age at this point.
*/
func StartInvasion(world *planet.World, maxAliens int) {
	initInvasion(world, maxAliens)
	invade(world)
}

func initInvasion(world *planet.World, maxAliens int) map[string]*Alien {
	nbAliens := maxAliens
	citiesWithAliens = make(map[string]*Alien, nbAliens)
	cityNames := world.CityNames()
	randomGenerator = initRandomGenerator()

	for i := 0; i < int(maxAliens); i++ {
		if len(cityNames) == 0 {
			println("All cities have been destroyed.")
			return citiesWithAliens
		}

		city, found := randomizeCity(world, cityNames)
		if !found {
			panic("Error initializing aliens")
		}

		alien := &Alien{id: i, city: city}
		if citiesWithAliens[city.Name()] != nil {
			alienInsideCity := citiesWithAliens[city.Name()]
			destroyCity(world, alienInsideCity, alien, city)
			//refresh city names since a city was removed
			cityNames = world.CityNames()
		} else {
			fmt.Printf("Setting alien %d to %s\n", alien.id, city.Name())
			citiesWithAliens[city.Name()] = alien
		}
	}

	return citiesWithAliens
}

func invade(world *planet.World) {
	fmt.Println("Starting invasion")

	for i := 0; i < 10000; i++ {
		for cityName, alien := range citiesWithAliens {

			if !alien.city.HasNeighbour() {
				fmt.Printf("Alien %d is trapped\n", alien.id)
				delete(citiesWithAliens, cityName)
				continue
			}

			city := alien.city
			goToCity := randomNeighbour(city)

			fmt.Printf("Moving alien %d to %s\n", alien.id, goToCity.Name())
			if citiesWithAliens[goToCity.Name()] != nil {
				alienInsideCity := citiesWithAliens[goToCity.Name()]
				destroyCity(world, alienInsideCity, alien, goToCity)
			} else {
				citiesWithAliens[goToCity.Name()] = alien
				alien.city = goToCity
			}
			delete(citiesWithAliens, cityName)
		}
	}
}

func randomizeCity(world *planet.World, cityNames []string) (*planet.City, bool) {
	var randIndex int
	if len(cityNames) == 0 {
		randIndex = 0
	} else {
		randIndex = randomGenerator.Intn(len(cityNames))
	}
	cityName := cityNames[randIndex]
	return world.GetCity(cityName)
}

func randomNeighbour(city *planet.City) *planet.City {
	availableDirections := city.NeighbourDirections()
	randIndex := randomGenerator.Intn(len(availableDirections))
	randDirection := city.NeighbourDirections()[randIndex]

	return city.GetNeighbour(randDirection)
}

func initRandomGenerator() *rand.Rand {
	randomSource := rand.NewSource(time.Now().UnixNano())
	return rand.New(randomSource)
}

func destroyCity(world *planet.World, alien1 *Alien, alien2 *Alien, city *planet.City) {
	delete(citiesWithAliens, city.Name())
	world.Remove(city.Name())
	city.RemoveCityFromNeighbour()
	fmt.Printf("%s has been destroyed by alien %d and alien %d!\n", city.Name(), alien1.id, alien2.id)
}
