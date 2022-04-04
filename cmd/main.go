package main

import (
	"fmt"
	"os"
	"strconv"

	"com.invasion/first/pkg/invasion"
	"com.invasion/first/pkg/io"
	"com.invasion/first/pkg/planet"
)

/*
	Main command of this application

	Usage : go run aliens.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS

	SOURCE_MAP_FILE_PATH should be in the same directory as the application
	NB_OF_ALIENS must be an positive integer

*/
func main() {
	sourceFilePath, nbAliens, err := parseArguments(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fsProvider := os.DirFS(".")
	lines, err := io.ReadLinesFromFile(fsProvider, sourceFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	world := planet.BuildWorld(lines)
	invasion.StartInvasion(world, nbAliens)

	fmt.Println("Here is the result map of the world after the invasion :")
	fmt.Print(world.ToString())
}

func parseArguments(args []string) (sourceFilePath string, nbAliens int, err error) {
	if len(args) != 3 {
		err = fmt.Errorf("Usage: go run aliens.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS")
		return
	}
	sourceFilePath = args[1]
	nbAliens, err = strconv.Atoi(args[2])
	if err != nil {
		err = fmt.Errorf("Error parsing command line argument. Usage: go run aliens.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS")
		return
	}
	return sourceFilePath, nbAliens, nil
}
