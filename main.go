package main

import (
	"fmt"
	"math/rand"
)

const (
	MapSize     = 50
	NumHouses   = 5
	NumFields   = 2
	GrassChance = 5
)

type TileInfo struct {
	Coordinates Coord
	IsWalkable  bool
}

type Coord struct {
	X int
	Y int
}

var (
	house = [][]int{
		{1, 5, 1, 5},
		{2, 6, 2, 6},
		{3, 7, 3, 7},
		{4, 8, 4, 8},
	}

	field = [][]int{
		{12, 15, 15, 15, 15, 16},
		{0, 18, 22, 22, 20, 13},
		{13, 22, 22, 22, 22, 13},
		{13, 22, 22, 22, 22, 13},
		{13, 19, 22, 22, 21, 13},
		{14, 15, 15, 15, 15, 17},
	}
)

func main() {
	// Initialize the map matrix with grass tiles (0)
	mapMatrix := make([][]int, MapSize)
	for i := 0; i < MapSize; i++ {
		mapMatrix[i] = make([]int, MapSize)
	}

	// Create 6 houses at random positions
	for i := 0; i < NumHouses; i++ {
		x := rand.Intn(MapSize - 4) // To ensure there's enough space for the house
		y := rand.Intn(MapSize - 4)
		createHouse(mapMatrix, x, y)
	}

	// Create 3 fields at random positions
	for i := 0; i < NumFields; i++ {
		x := rand.Intn(MapSize - 6) // To ensure there's enough space for the field
		y := rand.Intn(MapSize - 6)
		createField(mapMatrix, x, y)
	}

	// Cover the rest of the map with grass and flowers (5% of the time)
	for x := 0; x < MapSize; x++ {
		for y := 0; y < MapSize; y++ {
			if mapMatrix[x][y] == 0 {
				if rand.Intn(100) < GrassChance {
					mapMatrix[x][y] = rand.Intn(2) + 10 // Randomly choose between flowers (10) or grass (11)
				}
			}
		}
	}

	// Printing the mapMatrix (visual representation)
	for x := 0; x < MapSize; x++ {
		fmt.Printf("{")
		for y := 0; y < MapSize; y++ {
			fmt.Printf("%2d, ", mapMatrix[x][y])
		}
		fmt.Println("},")
	}
}

// Function to create a house at the specified position (x, y) on the map
func createHouse(mapMatrix [][]int, x, y int) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			mapMatrix[x+i][y+j] = house[i][j]
		}
	}
}

// Function to create a field at the specified position (x, y) on the map
func createField(mapMatrix [][]int, x, y int) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			mapMatrix[x+i][y+j] = field[i][j]
		}
	}
}
