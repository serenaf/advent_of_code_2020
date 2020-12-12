package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Customs struct {
	participants []string
	answers      map[rune]int
}

type CustomsList struct {
	customs []Customs
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func BuildCustomList(reader io.Reader) (CustomsList, error) {
	scanner := bufio.NewScanner(reader)
	customs := make([]Customs, 0)
	answers := make(map[rune]int)
	participants := make([]string, 0)
	for scanner.Scan() {
		l := scanner.Text()
		// this is one answer of a customer
		if len(l) != 0 {
			for _, v := range l {
				_, ok := answers[v]
				if ok {
					answers[v]++
				} else {
					answers[v] = 1
				}
			}
			participants = append(participants, l)
		}
		//this is one group of participants done
		if len(l) == 0 {
			newCustoms := Customs{
				participants: participants,
				answers:      answers,
			}
			customs = append(customs, newCustoms)
			answers = make(map[rune]int)
			participants = make([]string, 0)
		}
	}

	if len(participants) > 0 {
		newCustoms := Customs{
			participants: participants,
			answers:      answers,
		}
		customs = append(customs, newCustoms)
	}
	customsList := CustomsList{
		customs: customs,
	}
	return customsList, nil
}

func CountAnswers(c CustomsList) int {
	var answers int = 0
	for _, v := range c.customs {
		answers = answers + len(v.answers)
	}
	return answers
}

func CountAnswersAll(c CustomsList) int {
	var answers int = 0
	// v is a Custom struct
	for _, v := range c.customs {
		for _, item := range v.answers {
			if item == len(v.participants) {
				answers = answers + 1
			}
		}
	}
	return answers
}

func partOne() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	customsList, err := BuildCustomList(file)
	check(err)
	fmt.Println(CountAnswers(customsList))
}

func partTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	customsList, err := BuildCustomList(file)
	check(err)
	fmt.Println(CountAnswersAll(customsList))
}

func main() {
	partOne()
	partTwo()
}
