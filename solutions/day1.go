package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var list1 []int
var list2 []int

func GetInputFile() {
	file, err := os.Open("inputs/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		values := strings.Fields(line)
		if len(values) == 2 {
			// Convert the strings to integers
			val1, err1 := strconv.Atoi(values[0])
			val2, err2 := strconv.Atoi(values[1])
			if err1 == nil && err2 == nil {
				// Append to list1 and list2
				list1 = append(list1, val1)
				list2 = append(list2, val2)
			} else {
				fmt.Println("Error converting string to int:", err1, err2)
			}
		} else {
			fmt.Println("Unexpected line format:", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("List 1:", list1)
	fmt.Println("List 2:", list2)
}

func SortLists(list []int) {
	sort.Ints(list)
}

func ListsDistances(sortedList1 []int, sortedList2 []int) int {
	totalDistance := 0
	for i := 0; i < len(sortedList1); i++ {
		distance := sortedList1[i] - sortedList2[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	return totalDistance
}

func SolutionDay1() string {
	GetInputFile()

	SortLists(list1)
	SortLists(list2)

	totalDistance := ListsDistances(list1, list2)

	log.Printf("Sorted List 1: %v", list1)
	log.Printf("Sorted List 2: %v", list2)
	log.Printf("Total Distance: %d", totalDistance)
	result := fmt.Sprintf("The solution for Day 1 is: %d", totalDistance)

	return result
}
