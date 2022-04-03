package main

import (
	"fmt"
	"os"
	"strconv"

	"com.invasion/first/pkg/invasion"
	"com.invasion/first/pkg/io"
	"com.invasion/first/pkg/planet"
)

func main() {
	nbAliens, err := getNbAliens(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fsProvider := os.DirFS(".")
	lines, err := io.ReadLinesFromFile(fsProvider, "test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	world := planet.BuildWorld(lines)
	invasion.StartInvasion(world, nbAliens)

	fmt.Println("Here is the result map of the world after the invasion :")
	fmt.Print(world.ToString())
}

func getNbAliens(args []string) (nbAliens int, err error) {
	if len(args) != 2 {
		err = fmt.Errorf("Usage: go run aliens.go NB_OF_ALIENS")
		return
	}
	nbAliens, err = strconv.Atoi(args[1])
	if err != nil {
		err = fmt.Errorf("Error parsing command line argument. Usage: go run aliens.go NB_OF_ALIENS")
		return
	}
	return nbAliens, nil
}
