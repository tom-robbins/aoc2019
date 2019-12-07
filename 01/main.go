package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func massToFuel(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}

func massToFuelIncludingFuelMass(mass int) int {
	totalFuel := 0
	currentMass := mass
	for massToFuel(currentMass) > 0 {
		currentFuel := massToFuel(currentMass)
		totalFuel += currentFuel
		currentMass = currentFuel
	}
	return totalFuel
}

func main() {
	// open the intput file
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)

	// total fuel required
	totalFuel := 0

	// iterate through the input file and write to output file
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		totalFuel += massToFuelIncludingFuelMass(mass)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// write answer
	fmt.Printf("Solution: %v\n", strconv.Itoa(totalFuel))
}
