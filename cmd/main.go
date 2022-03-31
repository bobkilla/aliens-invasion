package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nbAliens, err := getNbAliens(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nbAliens)
}

func getNbAliens(args []string) (nbUfos int, err error) {
	if len(args) != 2 {
		err = fmt.Errorf("Usage: go run aliens.go NB_OF_ALIENS")
		return
	}
	nbUfos, err = strconv.Atoi(args[1])
	if err != nil {
		err = fmt.Errorf("Error parsing command line argument. Usage: go run aliens.go NB_OF_ALIENS")
		return
	}
	return nbUfos, nil
}
