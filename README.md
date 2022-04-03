# aliens-invasion

Simulation of aliens invading the Earth


## Assumptions
The program will :
- automatically add routes that are not specified. For example, if the city 'Foo' specifies a neighbour 'Bar' in the south, we will create the opposite route assuming the route is two-way.
- assumes that only one neighbour can exist in one direction. If the source map is inconsistent, the program will keep the first route specified and ignore any new route. (For example, 'Foo' has 'Bar' in the south, but 'Bar' specifies 'Baz' in the north, the program will keep the route between 'Foo' and 'Bar' and ignore the route between 'Bar' and 'Baz')
- assumes the map does not contain more than 10_000_000 cities to fit in the memory (this number depends on the free memory of the computer running the app. For this estimation, we assume having at least 10GB of free memory). If we want to handle more than that, we will use a database to be able to not load everything inside the memory.
- ignore any unspecified direction

The program does not use concurrency. I assumed that all aliens should move at the same pace.

The program will stop when :
- no more city exists in the world
- every alien is dead or trapped inside a city
- every alien has moved at least 10 000 times

## How to run the program
```
$ go run cmd/main.go SOURCE_MAP_FILE_PATH NB_OF_ALIENS
```
with :
`SOURCE_MAP_FILE_PATH`: location of the source map file on the disk
`NB_OF_ALIENS`: number of aliens invading the world

### Any questions?

Please send me an email, and I will answer quickly.

### License

Unlicense