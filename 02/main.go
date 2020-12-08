package main

import (
	"bufio"
	"fmt"
	"os"
)

type Password struct {
	minOccurence int
	maxOccurence int
	character    rune
	sequence     string
}

func (p Password) IsValid() bool {
	var cOccurence int = 0
	for _, c := range p.sequence {
		if c == p.character {
			cOccurence++
		}
	}
	return p.minOccurence <= cOccurence && cOccurence <= p.maxOccurence
}

func (p Password) IsTobogganValid() bool {
	position1 := rune(p.sequence[p.minOccurence-1]) == p.character
	position2 := rune(p.sequence[p.maxOccurence-1]) == p.character

	if position1 && position2 {
		return false
	}
	return position1 || position2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func PartOne() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var min, max int
	var character rune
	var sequence string
	var validCount int = 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &min, &max, &character, &sequence)
		p := Password{
			minOccurence: min,
			maxOccurence: max,
			character:    character,
			sequence:     sequence,
		}
		if p.IsValid() {
			validCount++
		}
	}
	fmt.Println("Valid Count", validCount)

}

func PartTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var min, max int
	var character rune
	var sequence string
	var validCount int = 0
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &min, &max, &character, &sequence)
		p := Password{
			minOccurence: min,
			maxOccurence: max,
			character:    character,
			sequence:     sequence,
		}
		if p.IsTobogganValid() {
			validCount++
		}
	}
	fmt.Println("Valid Count", validCount)
}

func main() {
	PartOne()
	PartTwo()
}
