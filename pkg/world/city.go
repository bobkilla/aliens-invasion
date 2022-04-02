package world

type City struct {
	name       string
	neighbours map[Direction]*City
	//aliens     []Alien
}
