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
	sourceFilePath, nbAliens, err := parseArguments(command)

	// then
	assert.Equal(sourceFilePath, "")
	assert.Equal(nbAliens, 0)
	assert.EqualError(err, "Usage: go run aliens.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS")
}

func Test_should_fail_when_too_many_arguments(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command", "10", "25", "32"}

	// when
	sourceFilePath, nbAliens, err := parseArguments(command)

	// then
	assert.Equal(sourceFilePath, "")
	assert.Equal(nbAliens, 0)
	assert.EqualError(err, "Usage: go run aliens.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS")
}

func Test_should_fail_when_string_argument_instead_of_int(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command", "test.txt", "str"}

	// when
	sourceFilePath, nbAliens, err := parseArguments(command)

	// then
	assert.Equal(sourceFilePath, "test.txt")
	assert.Equal(nbAliens, 0)
	assert.EqualError(err, "Error parsing command line argument. Usage: go run aliens.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS")
}

func Test_should_return_value(t *testing.T) {
	// given
	assert := assert.New(t)
	command := []string{"command", "test.txt", "40"}

	// when
	sourceFilePath, nbAliens, err := parseArguments(command)

	// then
	assert.Equal(sourceFilePath, "test.txt")
	assert.Equal(nbAliens, 40)
	assert.Nil(err)
}
