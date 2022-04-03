package invasion

import (
	"strconv"
	"testing"

	"com.invasion/first/pkg/planet"
	"github.com/stretchr/testify/assert"
)

func TestInitInvasion_with_one_city_one_alien(t *testing.T) {
	//given
	assert := assert.New(t)
	world, cities := givenWorld(1)

	//when
	result := initInvasion(world, 1)

	//then
	assert.Len(result, 1)
	assert.NotNil(result[cities[0].Name()].city)
}

func TestInitInvasion_should_destroy_city_with_one_city_and_two_aliens(t *testing.T) {
	//given
	assert := assert.New(t)
	world, _ := givenWorld(1)

	//when
	result := initInvasion(world, 2)

	//then
	assert.Len(result, 0)
	assert.True(world.IsEmpty())
}

func givenWorld(nbCities int) (*planet.World, []*planet.City) {
	world := planet.NewWorld()
	cities := make([]*planet.City, 0, nbCities)
	for i := 0; i < nbCities; i++ {
		city := planet.NewCity("city" + strconv.Itoa(i))
		world.Add(city)
		cities = append(cities, city)
	}

	return world, cities
}
