package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Position struct {
	x int
	y int
}

func getNextPos(previous Position, row string, xDir int, yDir int) Position {
	var newX int
	if previous.x+xDir < len(row) {
		newX = previous.x + xDir
	} else {
		newX = xDir - (len(row) - previous.x)
	}
	return Position{
		x: newX,
		y: previous.y + yDir,
	}
}

func calculateTrees(rows []string, xDir int, yDir int) int {
	var sumTrees int = 0
	var position = Position{
		x: 0,
		y: 0,
	}
	for i := range rows[0 : len(rows)-yDir] {
		position = getNextPos(position, rows[i], xDir, yDir)
		fmt.Println("x", position.x)
		fmt.Println("y", position.y)

		if position.y < len(rows) {
			row := rows[position.y]
			c := row[position.x]

			if c == '#' {
				sumTrees++
			}

		}
	}
	return sumTrees
}

func partOne() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var text []string
	var sumTrees int = 0
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	sumTrees = calculateTrees(text, 3, 1)
	fmt.Println("SumTrees", sumTrees)
}

func partTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	var sumTrees int = 0
	sumTrees = calculateTrees(text, 1, 1) * calculateTrees(text, 3, 1) * calculateTrees(text, 5, 1) * calculateTrees(text, 7, 1) * calculateTrees(text, 1, 2)
	// sumTrees = calculateTrees(text, 1, 2)
	fmt.Println("SumTrees", sumTrees)

}

func main() {
	// partOne()
	partTwo()
}
