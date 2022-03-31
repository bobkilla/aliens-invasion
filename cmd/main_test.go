package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_should_fail_when_no_argument(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command"}

	// when
	result, err := getNbAliens(command)

	// then
	assert.Equal(result, 0)
	assert.EqualError(err, "Usage: go run aliens.go NB_OF_ALIENS")
}

func Test_should_fail_when_too_many_arguments(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command", "10", "25"}

	// when
	result, err := getNbAliens(command)

	// then
	assert.Equal(result, 0)
	assert.EqualError(err, "Usage: go run aliens.go NB_OF_ALIENS")
}

func Test_should_fail_when_string_argument_instead_of_int(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command", "str"}

	// when
	result, err := getNbAliens(command)

	// then
	assert.Equal(result, 0)
	assert.EqualError(err, "Error parsing command line argument. Usage: go run aliens.go NB_OF_ALIENS")
}

func Test_should_return_value(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command", "40"}

	// when
	result, err := getNbAliens(command)

	// then
	assert.Equal(result, 40)
	assert.Nil(err)
}
