package io

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestFileReader_when_file_does_not_exists(t *testing.T) {
	//given
	assert := assert.New(t)
	fileName := "testFile.txt"
	fs := fstest.MapFS{}

	//when
	lines, err := readLinesFromFile(fs, fileName)

	//then
	assert.Nil(lines)
	assert.EqualError(err, "read file: open testFile.txt: file does not exist")
}

func TestFileReader(t *testing.T) {
	//given
	assert := assert.New(t)
	fileName := "testFile.txt"
	fs := fstest.MapFS{
		fileName: {
			Data: []byte("Foo north=Bar west=Baz south=Qu-ux\nBar south=Foo west=Bee"),
		},
	}

	//when
	lines, err := readLinesFromFile(fs, fileName)

	//then
	assert.Nil(err)
	assert.Len(lines, 2)
	assert.Equal(lines, []string{"Foo north=Bar west=Baz south=Qu-ux", "Bar south=Foo west=Bee"})
}
