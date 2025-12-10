package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TileArea struct {
	area     int
	tilePair []Tile
}

type Tile struct {
	X int
	Y int
}

// MAIN
func main() {
	// Slurp up the input
	inputFileContents, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	// get our datastructures ready
	redTiles := []Tile{} // it'll be a list of {x, y} coordinates

	// For each line in the input file
	lines := strings.Split(string(inputFileContents), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		vector := []int{}
		parts := strings.Split(strings.TrimSpace(line), ",")

		for _, posStr := range parts {
			position, _ := strconv.Atoi(posStr)
			vector = append(vector, position)
		}
		xPart, _ := strconv.Atoi(parts[0])
		yPart, _ := strconv.Atoi(parts[1])
		redTiles = append(redTiles, Tile{X: xPart, Y: yPart})
	}

	// build the graph edges
	areas := getAreas(redTiles)

	// PART 1
	fmt.Printf("Part 1 answer: %d\n", areas[0].area)

	// PART 2
	areas = getInternalAreas(redTiles)
	fmt.Printf("Part 2 answer: %d\n", areas[0].area)

}

func getAreas(redTiles []Tile) []TileArea {
	areas := []TileArea{}
	for tile1Index, _ := range redTiles {
		for tile2Index := tile1Index + 1; tile2Index < len(redTiles); tile2Index++ {
			tileArea := int(math.Abs(float64((redTiles[tile1Index].X - redTiles[tile2Index].X + 1) * (redTiles[tile1Index].Y - redTiles[tile2Index].Y + 1))))
			areas = append(areas, TileArea{area: tileArea, tilePair: []Tile{redTiles[tile1Index], redTiles[tile2Index]}})
		}
	}
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].area > areas[j].area
	})
	return areas
}

func getInternalAreas(redTiles []Tile) []TileArea {
	areas := []TileArea{}
	for tile1Index, _ := range redTiles {
		for tile2Index := tile1Index + 1; tile2Index < len(redTiles); tile2Index++ {
			intersecting := false
			for i, _ := range redTiles {
				if i == tile1Index || i == tile2Index {
					continue
				}
				if redTiles[tile1Index].X <= redTiles[tile2Index].X {
					if redTiles[tile1Index].Y <= redTiles[tile2Index].Y {
						if redTiles[tile1Index].X <= redTiles[i].X && redTiles[i].X <= redTiles[tile2Index].X && redTiles[tile1Index].Y <= redTiles[i].Y && redTiles[i].Y <= redTiles[tile2Index].Y {
							intersecting = true
							break
						}
					} else {
						if redTiles[tile1Index].X <= redTiles[i].X && redTiles[i].X <= redTiles[tile2Index].X && redTiles[tile1Index].Y >= redTiles[i].Y && redTiles[i].Y >= redTiles[tile2Index].Y {
							intersecting = true
							break
						}
					}
				} else {
					if redTiles[tile1Index].Y <= redTiles[tile2Index].Y {
						if redTiles[tile1Index].X >= redTiles[i].X && redTiles[i].X >= redTiles[tile2Index].X && redTiles[tile1Index].Y <= redTiles[i].Y && redTiles[i].Y <= redTiles[tile2Index].Y {
							intersecting = true
							break
						}
					} else {
						if redTiles[tile1Index].X >= redTiles[i].X && redTiles[i].X >= redTiles[tile2Index].X && redTiles[tile1Index].Y >= redTiles[i].Y && redTiles[i].Y >= redTiles[tile2Index].Y {
							intersecting = true
							break
						}
					}
				}
				if i == (len(redTiles) - 1) {
					continue
				}
				intersecting = checkSquareIntersection(redTiles[tile1Index], redTiles[tile2Index], redTiles[i], redTiles[i+1])
				if intersecting {
					break
				}
			}
			if !intersecting {
				tileArea := int(math.Abs(float64((redTiles[tile1Index].X - redTiles[tile2Index].X + 1) * (redTiles[tile1Index].Y - redTiles[tile2Index].Y + 1))))
				areas = append(areas, TileArea{area: tileArea, tilePair: []Tile{redTiles[tile1Index], redTiles[tile2Index]}})
			}
		}
	}
	sort.Slice(areas, func(i, j int) bool {
		return areas[i].area > areas[j].area
	})
	return areas
}

func checkSquareIntersection(corner1 Tile, corner2 Tile, testPoint1 Tile, testPoint2 Tile) bool {
	if checkSegmentIntersection(corner1, Tile{corner2.X, corner1.Y}, testPoint1, testPoint2) {
		return true
	}
	if checkSegmentIntersection(Tile{corner2.X, corner1.Y}, corner2, testPoint1, testPoint2) {
		return true
	}
	if checkSegmentIntersection(corner2, Tile{corner1.X, corner2.Y}, testPoint1, testPoint2) {
		return true
	}
	if checkSegmentIntersection(Tile{corner1.X, corner2.Y}, corner2, testPoint1, testPoint2) {
		return true
	}
	return false
}

func checkSegmentIntersection(p1 Tile, p2 Tile, q1 Tile, q2 Tile) bool {
	o1 := orientation(p1, p2, q1)
	o2 := orientation(p1, p2, q2)
	o3 := orientation(p1, q1, q2)
	o4 := orientation(p2, q1, q2)

	// General (non-collinear) case
	if o1 != o2 && o3 != o4 {
		return true
	}
	return false
}

func orientation(p, q, r Tile) bool {
	return (r.Y-p.Y)*(q.X-p.X) > (q.Y-p.Y)*(r.X-p.X)
}
