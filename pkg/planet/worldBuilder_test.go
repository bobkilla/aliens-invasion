package planet

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkAliens(b *testing.B) {
	cities := make(map[string]*City)
	for i := 0; i < b.N; i++ {
		name := strconv.Itoa(i)
		north := &City{name: name + "2", neighbours: make(map[Direction]*City)}
		south := &City{name: name + "3", neighbours: make(map[Direction]*City)}
		east := &City{name: name + "4", neighbours: make(map[Direction]*City)}
		west := &City{name: name + "5", neighbours: make(map[Direction]*City)}
		cities[name] = &City{name: name, neighbours: map[Direction]*City{North: north, South: south, East: east, West: west}}
	}
}

func TestWorldBuilder_when_lines_is_empty(t *testing.T) {
	//given
	assert := assert.New(t)
	lines := []string{}

	//when
	world := BuildWorld(lines)

	//then
	assert.NotNil(world)
	assert.Empty(world.cities)
}

func TestWorldBuilder_when_line_does_not_contain_any_space(t *testing.T) {
	//given
	assert := assert.New(t)
	lines := []string{"city"}

	//when
	world := BuildWorld(lines)

	//then
	assert.NotNil(world)
	assert.Len(world.cities, 1)
	assertCityHasNeighbours(assert, world, "city", map[Direction]string{})
}

func TestWorldBuilder_when_line_has_one_neighbour_with_multiple_equal_char(t *testing.T) {
	//given
	assert := assert.New(t)
	lines := []string{"city north=north=cityName"}

	//when
	world := BuildWorld(lines)

	//then
	assert.NotNil(world)
	assert.Len(world.cities, 1)
	assertCityHasNeighbours(assert, world, "city", map[Direction]string{})
}

func TestWorldBuilder_when_line_has_one_neighbour_with_unkown_direction(t *testing.T) {
	//given
	assert := assert.New(t)
	lines := []string{"city whatever=otherCity"}

	//when
	world := BuildWorld(lines)

	//then
	assert.NotNil(world)
	assert.Len(world.cities, 1)
	assertCityHasNeighbours(assert, world, "city", map[Direction]string{})
	assertCityNotExists(assert, world, "otherCity")
}

func TestWorldBuilder_should_ignore_route_when_a_route_was_already_specified_for_this_direction(t *testing.T) {
	//given
	assert := assert.New(t)
	lines := []string{"Foo north=Bar", "Baz south=Foo"}

	//when
	world := BuildWorld(lines)

	//then
	assert.NotNil(world)
	assert.Len(world.cities, 3)
	assertCityHasNeighbours(assert, world, "Foo", map[Direction]string{North: "Bar"})
	assertCityHasNeighbours(assert, world, "Bar", map[Direction]string{South: "Foo"})
	assertCityHasNeighbours(assert, world, "Baz", map[Direction]string{})
}

func TestWorldBuilder_when_format_is_correct(t *testing.T) {
	//given
	assert := assert.New(t)
	lines := []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"}

	//when
	world := BuildWorld(lines)

	//then
	assert.NotNil(world)
	assert.Len(world.cities, 5)
	assertCityHasNeighbours(assert, world, "Foo", map[Direction]string{North: "Bar", West: "Baz", South: "Qu-ux"})
	assertCityHasNeighbours(assert, world, "Bar", map[Direction]string{South: "Foo", West: "Bee"})
	assertCityHasNeighbours(assert, world, "Baz", map[Direction]string{East: "Foo"})
	assertCityHasNeighbours(assert, world, "Qu-ux", map[Direction]string{North: "Foo"})
	assertCityHasNeighbours(assert, world, "Bee", map[Direction]string{East: "Bar"})
}

func assertCityHasNeighbours(assert *assert.Assertions, world *World, cityName string, neighbours map[Direction]string) {
	city, found := world.GetCity(cityName)
	assert.True(found)
	assert.NotNil(city)

	for direction, neighbourName := range neighbours {
		neighbour, found := world.GetCity(neighbourName)
		assert.True(found)
		assert.NotNil(neighbour)
		assert.Equal(city.GetNeighbour(direction), neighbour)
	}
}

func assertCityNotExists(assert *assert.Assertions, world *World, cityName string) {
	city, found := world.GetCity(cityName)
	assert.False(found)
	assert.Nil(city)
}
