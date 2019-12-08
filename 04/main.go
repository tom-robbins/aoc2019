package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func validate(pass int) bool {
	passString := strconv.Itoa(pass)

	lastDigit := -1
	duplicateGroupLengths := make(map[int]int)
	for _, c := range passString {
		digit := int(c - '0')

		// password digits must be increasing
		if digit < lastDigit {
			return false
		}

		if digit == lastDigit {
			_, ok := duplicateGroupLengths[digit]
			if ok {
				duplicateGroupLengths[digit]++
			} else {
				duplicateGroupLengths[digit] = 2
			}
		}
		lastDigit = digit
	}

	// // Part 1
	// containsGroupOfAtLeastTwo := false
	// for _, groupLength := range duplicateGroupLengths {
	// 	if groupLength >= 2 {
	// 		containsGroupOfAtLeastTwo = true
	// 	}
	// }
	// return containsGroupOfAtLeastTwo

	// Part 2
	containsGroupOfExactlyTwo := false
	for _, groupLength := range duplicateGroupLengths {
		if groupLength == 2 {
			containsGroupOfExactlyTwo = true
		}
	}
	return containsGroupOfExactlyTwo
}

func main() {
	// open the intput file
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()

	// read IntCode program into programCopy
	inputBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	inputString := strings.TrimSuffix(string(inputBytes), "\n")
	inputStringArray := strings.Split(inputString, "-")

	minPass, err := strconv.Atoi(inputStringArray[0])
	if err != nil {
		log.Fatal(err)
	}
	maxPass, err := strconv.Atoi(inputStringArray[1])
	if err != nil {
		log.Fatal(err)
	}

	validPasswords := 0
	for i := minPass; i < maxPass; i++ {
		if validate(i) {
			validPasswords++
		}
	}
	fmt.Printf("%v\n", validPasswords)
}
