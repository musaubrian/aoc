package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./strat_guide")
	if err != nil {
		log.Fatal("ERROR:", err)
	}
	defer file.Close()

	fScanner := bufio.NewScanner(file)

	var score int
	scores := map[string]int{
		"B X": 1, "C Y": 2,
		"A Z": 3, "A X": 4,
		"B Y": 5, "C Z": 6,
		"C X": 7, "A Y": 8,
		"B Z": 9,
	}

	for fScanner.Scan() {
		score += scores[fScanner.Text()]
	}
	fmt.Println(score)
}
