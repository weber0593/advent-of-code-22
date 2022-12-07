package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputData, err := ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	chunkedInputs := ChunkInputsByNewline(inputData)

	var sortedCalorieCounts []int

	for _, singleElfData := range chunkedInputs {
		calories := TotalCalories(singleElfData)
		sortedCalorieCounts = InsertIntoSlice(sortedCalorieCounts, calories)
	}
	fmt.Printf("The top 3 elfs with the most Calories have:\n %v\n %v\n %v\nFor a combined: %v\n",
		sortedCalorieCounts[len(sortedCalorieCounts)-1],
		sortedCalorieCounts[len(sortedCalorieCounts)-2],
		sortedCalorieCounts[len(sortedCalorieCounts)-3],
		sortedCalorieCounts[len(sortedCalorieCounts)-1]+sortedCalorieCounts[len(sortedCalorieCounts)-2]+sortedCalorieCounts[len(sortedCalorieCounts)-3],
	)
}

func InsertIntoSlice(sortedSlice []int, value int) []int {
	if len(sortedSlice) == 0 { //base case, slice is empty, just return the new value as a slice
		return []int{value}
	}
	if value < sortedSlice[0] { //if the value is lower than the lowest element, insert at the front
		return append([]int{value}, sortedSlice...)
	}
	if value >= sortedSlice[len(sortedSlice)-1] { //if the value is greater than the highest element, append to the end
		return append(sortedSlice, value)
	}
	//for everything else, we have to find the proper place
	for i := 0; i < len(sortedSlice); i++ {
		if sortedSlice[i] <= value && sortedSlice[i+1] > value {
			//this looks a little funny but is some magic to "move everything in the slice right one position"
			sortedSlice = append(sortedSlice[:i+2], sortedSlice[i+1:]...)
			//then we can insert our value where we want
			sortedSlice[i+1] = value
			return sortedSlice
		}
	}
	fmt.Printf("Uhhh.. we shouldnt get here.. %v, %v\n", sortedSlice, value)
	return sortedSlice
}

func ChunkInputsByNewline(inputs []string) [][]string {
	var chunkedInputs [][]string
	var singleChunk []string
	for _, item := range inputs {
		if len(item) == 0 { //if we hit a newline, then we are done with this chunk
			if len(singleChunk) > 0 {
				chunkedInputs = append(chunkedInputs, singleChunk)
				singleChunk = []string{}
				continue
			}
		}
		singleChunk = append(singleChunk, item)
	}
	if len(singleChunk) > 0 { //once we are done, make sure we add the last chunk
		chunkedInputs = append(chunkedInputs, singleChunk)
		singleChunk = []string{}
	}
	return chunkedInputs
}

func TotalCalories(items []string) int {
	totalCalories := 0
	for _, item := range items {
		value, err := strconv.Atoi(item)
		if err != nil {
			fmt.Printf("Invalid input row: %s", item)
			continue
		}
		totalCalories += value
	}
	return totalCalories
}

func ReadFile(Path string) ([]string, error) {
	var data []string
	f, err := os.Open(Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
