package main

import (
	calories "advent-of-code-22/internal/calorie-counting"
	rps "advent-of-code-22/internal/rock-paper-scissors"
	"flag"
	"fmt"
	"os"
)

func main() {
	puzzleNum := flag.Int("puzzle", 0, "Number of the puzzle to run")

	flag.Parse()
	if *puzzleNum == 0 {
		fmt.Println("Please enter the number of the puzzle you wish to run")
		os.Exit(0)
	}
	switch *puzzleNum {
	case 1:
		calories.Run("internal/calorie-counting/input.txt") //TODO: fix this path maybe
	case 2:
		rps.Run("internal/rock-paper-scissors/input.txt")
	default:
		fmt.Printf("Unknown puzzle number %d\n", *puzzleNum)
	}
}
