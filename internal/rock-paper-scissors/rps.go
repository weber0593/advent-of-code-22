package rock_paper_scissors

import (
	file "advent-of-code-22/pkg/file-reader"
	"fmt"
	"os"
)

func Run(inputFilePath string) {
	inputData, err := file.ReadFile(inputFilePath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	totalScore := 0
	for _, round := range inputData {
		score, err := CalcScoreSingleRound(round)
		if err != nil {
			fmt.Printf("Error with round: %s\n", err)
			continue
		}
		totalScore += score
	}
	fmt.Printf("Total Score: %d\n", totalScore)
}

func CalcScoreSingleRound(input string) (int, error) {
	if len(input) != 3 {
		return -1, fmt.Errorf("invalid input string: %s", input)
	}
	totalScore := 0
	opponentMove := rune(input[0])
	shouldWin := rune(input[2])
	var yourMove rune
	switch shouldWin {
	case 'X':
		switch opponentMove {
		case 'A':
			yourMove = 'Z'
		case 'B':
			yourMove = 'X'
		case 'C':
			yourMove = 'Y'
		}
	case 'Y':
		switch opponentMove {
		case 'A':
			yourMove = 'X'
		case 'B':
			yourMove = 'Y'
		case 'C':
			yourMove = 'Z'
		}
	case 'Z':
		switch opponentMove {
		case 'A':
			yourMove = 'Y'
		case 'B':
			yourMove = 'Z'
		case 'C':
			yourMove = 'X'
		}
	}
	totalScore += CalcScoreForMove(yourMove)
	totalScore += CalcRoundWinLossScore(opponentMove, yourMove)
	return totalScore, nil
}

func CalcRoundWinLossScore(opponentMove rune, yourMove rune) int {
	switch opponentMove {
	case 'A':
		switch yourMove {
		case 'X':
			return 3
		case 'Y':
			return 6
		case 'Z':
			return 0
		}
	case 'B':
		switch yourMove {
		case 'X':
			return 0
		case 'Y':
			return 3
		case 'Z':
			return 6
		}
	case 'C':
		switch yourMove {
		case 'X':
			return 6
		case 'Y':
			return 0
		case 'Z':
			return 3
		}
	}
	fmt.Printf("Invalid Matchup %v, %v\n", opponentMove, yourMove)
	return 0
}

func CalcScoreForMove(input rune) int {
	switch input {
	case 'X':
		return 1
	case 'Y':
		return 2
	case 'Z':
		return 3
	default:
		return 0
	}
}
