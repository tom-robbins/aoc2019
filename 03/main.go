package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type wirePathSegment struct {
	direction string
	magnitude int
}

type coordinate struct {
	x int
	y int
}

func manhattanDistance(p coordinate, q coordinate) int {
	return int(math.Abs(float64(p.x)-float64(q.x)) + math.Abs(float64(p.y)-float64(q.y)))
}

func getCoordinatesForWirePath(wirePath []wirePathSegment) []coordinate {
	var coordinates []coordinate
	currentCoordinate := coordinate{0, 0}
	for _, segment := range wirePath {
		for i := 0; i < segment.magnitude; i++ {
			var newCoordinate coordinate
			switch segment.direction {
			case "R":
				newCoordinate = coordinate{currentCoordinate.x + 1, currentCoordinate.y}
			case "U":
				newCoordinate = coordinate{currentCoordinate.x, currentCoordinate.y + 1}
			case "L":
				newCoordinate = coordinate{currentCoordinate.x - 1, currentCoordinate.y}
			case "D":
				newCoordinate = coordinate{currentCoordinate.x, currentCoordinate.y - 1}
			}
			coordinates = append(coordinates, newCoordinate)
			currentCoordinate = newCoordinate
		}
	}
	return coordinates
}

func intersect(a map[coordinate]int, b map[coordinate]int) map[coordinate]int {
	coords := make(map[coordinate]int)
	for aCoord, aSteps := range a {
		bSteps, ok := b[aCoord]
		if ok {
			coords[aCoord] = aSteps + bSteps
		}
	}
	return coords
}

func getWireCoordinateIntersections(wirePathsCoordinates [][]coordinate) map[coordinate]int {
	var wireCoordinateMaps []map[coordinate]int

	// use map as a set implementation
	for _, wireCoordinates := range wirePathsCoordinates {
		newMap := make(map[coordinate]int)
		for i, coord := range wireCoordinates {
			_, ok := newMap[coord]
			if ok {
				continue
			} else {
				newMap[coord] = i + 1
			}
		}
		wireCoordinateMaps = append(wireCoordinateMaps, newMap)
	}

	// get intersection of the two maps
	intersections := intersect(wireCoordinateMaps[0], wireCoordinateMaps[1])
	return intersections
}

func main() {
	// open the intput file
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()

	// read IntCode program into programCopy
	scanner := bufio.NewScanner(inputFile)

	var wires []string
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// generate wire paths representation
	var wirePaths [][]wirePathSegment
	for i, wire := range wires {
		var newSlice []wirePathSegment
		wirePaths = append(wirePaths, newSlice)
		for _, pathString := range strings.Split(wire, ",") {
			direction := pathString[0]
			magnitude, err := strconv.Atoi(pathString[1:])
			if err != nil {
				log.Fatal(err)
			}
			wirePaths[i] = append(wirePaths[i], wirePathSegment{string(direction), magnitude})
		}
	}

	// get grid spaces occupied by each wire
	var wireCoordinates [][]coordinate
	for _, wirePath := range wirePaths {
		wireCoordinates = append(wireCoordinates, getCoordinatesForWirePath(wirePath))
	}

	// get intersections
	intersections := getWireCoordinateIntersections(wireCoordinates)

	// get closest intersection
	minDistance := math.MaxInt32
	var minIntersection coordinate

	// // Manhattan Distance
	// for intersection := range intersections {
	// 	distance := manhattanDistance(coordinate{0, 0}, intersection)
	// 	if distance < minDistance {
	// 		minDistance = distance
	// 		minIntersection = intersection
	// 	}
	// }

	// Step Distance
	for intersection, steps := range intersections {
		if steps < minDistance {
			minDistance = steps
			minIntersection = intersection
		}
	}
	fmt.Printf("Coordinates: %v, Distance: %v", minIntersection, minDistance)
}
