package planet

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitInvasion_with_one_city_one_alien(t *testing.T) {
	//given
	assert := assert.New(t)
	world, cities := givenWorld(1)

	//when
	result := InitInvasion(world, 1)

	//then
	assert.Len(result, 1)
	assert.Equal(result[0].city, cities[0])
}

func TestInitInvasion_should_destroy_city_with_one_city_and_two_aliens(t *testing.T) {
	//given
	assert := assert.New(t)
	world, _ := givenWorld(1)

	//when
	result := InitInvasion(world, 2)

	//then
	assert.Len(result, 0)
	assert.Empty(world.cities)
}

func TestDestroyCity(t *testing.T) {
	//given
	assert := assert.New(t)
	world, cities := givenWorld(1)
	alien1 := Alien{id: 1}
	alien2 := Alien{id: 2}
	aliens := map[int]*Alien{1: &alien1, 2: &alien2}

	//when
	destroyCity(world, aliens, &alien1, &alien2, cities[0])

	//then
	assert.Empty(aliens)
	assert.Empty(world.cities)
}

func givenWorld(nbCities int) (*World, []*City) {
	world := newWorld()
	cities := make([]*City, 0, nbCities)
	for i := 0; i < nbCities; i++ {
		city := newCity("city" + strconv.Itoa(i))
		world.add(city)
		cities = append(cities, city)
	}

	return world, cities
}
