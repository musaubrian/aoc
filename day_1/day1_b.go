package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	maxCalsA := 0
	maxCalsB := 0
	maxCalsC := 0
	currentElfCals := 0

	file, err := os.Open("./puzzle")
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
	defer file.Close()

	fScanner := bufio.NewScanner(file)
	for fScanner.Scan() {
		cals, err := strconv.Atoi(fScanner.Text())
		currentElfCals += cals

		// Found an empty line
		if err != nil {
			if currentElfCals > maxCalsC {
				maxCalsC = currentElfCals
			}
			if maxCalsC > maxCalsB {
				maxCalsC, maxCalsB = maxCalsB, maxCalsC
			}
			if maxCalsB > maxCalsA {
				maxCalsB, maxCalsA = maxCalsA, maxCalsB
			}

			currentElfCals = 0
		}
	}

	fmt.Println(maxCalsA + maxCalsB + maxCalsC)
}
