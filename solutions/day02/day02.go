package day02

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ===> Day 2, Part 1 <===
func Part1(inputfile string) {
	var totalSafe int = 0
	inputs, err := loadInputs(inputfile)
	if err != nil {
		fmt.Printf("Error when loading inputs: %s\n", err)
	}
	for _, slice := range *inputs {
		var diff bool = false
		sorted := isSorted(&slice)
		if sorted {
			diff = checkDiffer(&slice)
			if diff {
				totalSafe++
			}
		}
	}
	fmt.Printf("Total: %d\n", totalSafe)
}

// ===> Day 2, Part 2 <===
func Part2(inputfile string) {
	var totalSafe int = 0
	inputs, err := loadInputs(inputfile)
	if err != nil {
		fmt.Printf("Error when loading inputs: %s\n", err)
	}
	for _, slice := range *inputs {
		var diff bool = false
		sorted := isSorted(&slice)
		if sorted {
			diff = checkDiffer(&slice)
			if diff {
				// all good
				totalSafe++
			} else {
				// sorted, but diff not between 1 and 3
				for x := 0; x < len(slice); x++ {
					newslice := removeItemFromSlice(&slice, x)
					if checkDiffer(newslice) {
						totalSafe++
						break
					}
				}
			}
		} else {
			// not sorted, need to repeat both checks
			for x := 0; x < len(slice); x++ {
				newslice := removeItemFromSlice(&slice, x)
				if !isSorted(newslice) {
					continue
				}
				if checkDiffer(newslice) {
					totalSafe++
					break
				}
			}
		}
	}
	fmt.Printf("Total: %d\n", totalSafe)
}

func loadInputs(filepath string) (*[][]int, error) {
	inputs, err := os.Open(filepath)
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %d", err))
	}
	defer inputs.Close()

	scanner := bufio.NewScanner(inputs)

	var result [][]int
	for scanner.Scan() {
		line := scanner.Text()

		strSlice := strings.Split(line, " ")
		var intSlice []int
		for _, char := range strSlice {
			num, err := strconv.Atoi(char)
			if err == nil {
				intSlice = append(intSlice, num)
			}
		}

		result = append(result, intSlice)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s", err)
		return nil, err
	}

	return &result, nil
}

func isSorted(slice *[]int) bool {
	return sort.SliceIsSorted(*slice, func(i, j int) bool {
		return (*slice)[i] < (*slice)[j]
	}) || sort.SliceIsSorted(*slice, func(i, j int) bool {
		return (*slice)[i] > (*slice)[j]
	})
}

func checkDiffer(slice *[]int) bool {
	for i := 0; i < len(*slice)-1; i++ {
		differ := (*slice)[i] - (*slice)[i+1]
		absDiffer := int(math.Abs(float64(differ)))
		if absDiffer < 1 || absDiffer > 3 {
			return false
		}
	}
	return true
}

func removeItemFromSlice(slice *[]int, index int) *[]int {
	var result []int
	for i := 0; i < len(*slice); i++ {
		if i == index {
			continue
		}
		result = append(result, (*slice)[i])
	}
	return &result
}
