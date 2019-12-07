package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func intsToStrings(arr []int) []string {
	resultArr := make([]string, len(arr))
	for i, v := range arr {
		newValue := strconv.Itoa(v)
		resultArr[i] = newValue
	}
	return resultArr
}

func stringsToInts(arr []string) []int {
	resultArr := make([]int, len(arr))
	for i, v := range arr {
		newValue, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		resultArr[i] = newValue
	}
	return resultArr
}

func trial(program []int, noun int, verb int) (int, error) {
	programCopy := make([]int, len(program))
	copy(programCopy, program)

	programCopy[1] = noun
	programCopy[2] = verb

	opIndex := 0
	for {
		opCode := programCopy[opIndex]
		switch opCode {
		case 1:
			programCopy[programCopy[opIndex+3]] = programCopy[programCopy[opIndex+1]] + programCopy[programCopy[opIndex+2]]
			opIndex += 4
		case 2:
			programCopy[programCopy[opIndex+3]] = programCopy[programCopy[opIndex+1]] * programCopy[programCopy[opIndex+2]]

			opIndex += 4
		case 99:
			return programCopy[0], nil
		default:
			return 0, errors.New("Bad Op code")
		}
	}
}

func main() {
	// open the intput file
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()

	// read IntCode program into programCopy
	intCodeProgramBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	intCodeProgramString := strings.TrimSuffix(string(intCodeProgramBytes), "\n")
	intCodeProgramStringArray := strings.Split(intCodeProgramString, ",")
	intCodeProgramIntArray := stringsToInts(intCodeProgramStringArray)

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			// fmt.Printf("%v %v\n", i, j)
			result, err := trial(intCodeProgramIntArray, i, j)
			if err != nil {
				log.Fatal(err)
			}
			if result == 19690720 {
				fmt.Printf("Correct noun, verb: %d, %d\n", i, j)
				fmt.Printf("Solution: %d\n", 100*i+j)
			}
		}
	}
}
