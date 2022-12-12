package main

import (
	calories "advent-of-code-22/calorie-counting"
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
		calories.Run()
	case 2:
		//TODO:
	default:
		fmt.Printf("Unknown puzzle number %d\n", *puzzleNum)
	}
}
