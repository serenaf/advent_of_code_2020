package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type TreeMap struct {
	lines []string
}

func NewTreeMap(reader io.Reader) (TreeMap, error) {
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return TreeMap{lines: lines}, nil
}

func (t TreeMap) CountTrees(column, row, right, down int) int {
	trees := 0
	for row < len(t.lines) {
		if t.lines[row][column] == '#' {
			trees++
		}

		column = (column + right) % len(t.lines[row])
		row = row + down
	}
	return trees
}

func partOne() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	treeMap, error := NewTreeMap(file)
	check(error)
	fmt.Println("SumTrees", treeMap.CountTrees(0, 0, 3, 1))
}

func partTwo() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	treeMap, error := NewTreeMap(file)
	check(error)
	fmt.Println("SumTrees", treeMap.CountTrees(0, 0, 3, 1)*treeMap.CountTrees(0, 0, 5, 1)*treeMap.CountTrees(0, 0, 7, 1)*treeMap.CountTrees(0, 0, 1, 2))

}

func main() {
	partOne()
	partTwo()
}
