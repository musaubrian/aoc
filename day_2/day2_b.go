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
		"B X": 1, "C X": 2,
		"A X": 3, "A Y": 4,
		"B Y": 5, "C Y": 6,
		"C Z": 7, "A Z": 8,
		"B Z": 9,
	}

	for fScanner.Scan() {
		score += scores[fScanner.Text()]
	}
	fmt.Println(score)
}
