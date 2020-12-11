package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type BoardingPass struct {
	row    int
	column int
	seatId int
}

func FindBoardingPass(input string) BoardingPass {
	rowStart := 0
	rowEnd := 128
	columnStart := 0
	columnEnd := 8
	for _, s := range input {
		switch s {
		case 'B':
			rowStart = rowStart + (rowEnd-rowStart)/2
		case 'F':
			rowEnd = rowStart + (rowEnd-rowStart)/2
		case 'R':
			columnStart = columnStart + (columnEnd-columnStart)/2
		case 'L':
			columnEnd = columnStart + (columnEnd-columnStart)/2
		}

	}
	bp := BoardingPass{
		row:    rowStart,
		column: columnStart,
		seatId: rowStart*8 + columnStart,
	}
	return bp
}

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
	seatId := 0
	for scanner.Scan() {
		l := scanner.Text()
		bp := FindBoardingPass(l)
		if bp.seatId > seatId {
			seatId = bp.seatId
		}
	}
	fmt.Println("the highest seatid is", seatId)
}

func partTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var seatIdList []int
	for scanner.Scan() {
		l := scanner.Text()
		bp := FindBoardingPass(l)
		seatIdList = append(seatIdList, bp.seatId)
	}
	sort.Ints(seatIdList)
	for i := 0; i < len(seatIdList)-1; i++ {
		if seatIdList[i+1] != seatIdList[i]+1 {
			fmt.Println("Found a space", i, i+1, seatIdList[i], seatIdList[i+1])
		}

	}
}

func main() {
	partOne()
	partTwo()
}
