package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var text []string
	var candidateA, candidateB int
	var numbers = make(map[int]bool)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	// build up the hashmap consisting of the numbers
	for _, number := range text {
		iNumber, _ := strconv.Atoi(number)
		numbers[iNumber] = true
	}

	for key, _ := range numbers {
		candidateA = key
		candidateB = 2020 - candidateA
		_, ok := numbers[candidateB]
		if ok {
			break
		}
	}
	fmt.Println("Candidate A", candidateA)
	fmt.Println("Candidate B", candidateB)
	fmt.Println("End Result", candidateA*candidateB)
}

func partTwo() int {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var text []string
	var numbers []int
	var seenNumbers = make(map[int]bool)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	for _, line := range text {
		element, _ := strconv.Atoi(line)
		numbers = append(numbers, element)
		seenNumbers[element] = true
	}

	for i, candidateA := range numbers {
		for _, candidateB := range numbers[i+1:] {
			candidateC := 2020 - (candidateA + candidateB)
			_, ok := seenNumbers[candidateC]
			if ok {
				fmt.Println("Candidate A", candidateA)
				fmt.Println("Candidate B", candidateB)
				fmt.Println("Candidate C", candidateC)
				return candidateA * candidateB * candidateC
			}
		}
	}
	return 0
}

func main() {
	partOne()
	result := partTwo()
	fmt.Println("End Result", result)
}
