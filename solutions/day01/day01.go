package day01

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
)

// ===> Day 1, Part 1 <===
func Part1(inputs string) {
	leftList, rightList, err := loadInputs(inputs)
	if err != nil {
		panic("Unable to scan inputs")
	}
	slices.Sort(*leftList)
	slices.Sort(*rightList)

	var results []int
	for x := 0; x < len(*leftList); x++ {
		newDistance := calculateDistances((*leftList)[x], (*rightList)[x])
		results = append(results, newDistance)
	}

	var total int = 0
	for i := 0; i < len(results); i++ {
		total = total + results[i]
	}

	fmt.Printf("Result: %d\n", total)
}

// ===> Day 1, Part 2 <===
func Part2(inputs string) {
	leftList, rightList, err := loadInputs(inputs)
	if err != nil {
		panic("Unable to scan inputs")
	}
	slices.Sort(*leftList)
	slices.Sort(*rightList)

	results := calculateOcurrences(leftList, rightList)

	var total int = 0
	for i := 0; i < len(*results); i++ {
		total = total + (*results)[i]
	}

	fmt.Printf("Result: %d\n", total)
}

func loadInputs(filepath string) (*[]int, *[]int, error) {
	var leftList []int
	var rightList []int

	inputs, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %d", err))
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	for scanner.Scan() {
		line := scanner.Text()
		var leftnum, rightnum int

		_, err := fmt.Sscanf(line, "%d %d", &leftnum, &rightnum)
		if err != nil {
			msg := fmt.Sprintf("Error parsing line: %s", line)
			return nil, nil, errors.New(msg)
		}
		leftList = append(leftList, leftnum)
		rightList = append(rightList, rightnum)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		return nil, nil, err
	}

	return &leftList, &rightList, nil
}

func calculateDistances(a, b int) int {
	distance := math.Abs(float64(a - b))
	return int(distance)
}

func calculateOcurrences(leftList, rightList *[]int) *[]int {
	var results []int
	for x := 0; x < len(*leftList); x++ {
		numL := (*leftList)[x]
		count := 0
		for y := 0; y < len(*rightList); y++ {
			numR := (*rightList)[y]
			if numL == numR {
				count++
			}
			if numL < numR {
				continue
			}
		}
		results = append(results, numL*count)
	}
	return &results
}
