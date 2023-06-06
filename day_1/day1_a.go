package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	maxCals := 0
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
			if currentElfCals > maxCals {
				maxCals = currentElfCals
			}
			currentElfCals = 0
		}
	}

	fmt.Println(maxCals)
}
