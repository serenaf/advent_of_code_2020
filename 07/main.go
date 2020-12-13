package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne() int {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	contains := make(map[string]map[string]int)
	shinyGoldParents := 0
	for scanner.Scan() {
		l := scanner.Text()
		subStrings := strings.Split(l, " bags contain ")
		if subStrings[1] == "no other bags." {
			continue
		}
		bagType := subStrings[0]
		contains[bagType] = make(map[string]int)
		contents := strings.Split(subStrings[1], ", ")

		for _, v := range contents {
			v = strings.TrimSuffix(v, " bags")
			v = strings.TrimSuffix(v, " bag")
			subParts := strings.SplitN(v, " ", 2)
			num, err := strconv.Atoi(subParts[0])
			content := subParts[1]
			if err != nil {
				fmt.Printf("Failed to parse %s\n", l)
				break
			}
			contains[bagType][content] = num
		}
	}

	itemsToLookFor := []string{"shiny gold"}
	found := make(map[string]bool)
	for {
		// first we are looking for shiny gold parent
		for _, item := range itemsToLookFor {
			// loop through the contains hashmap (string -> bagDetail (String, int))
			for key, value := range contains {
				for items := range value {
					if strings.Contains(items, item) {
						if !found[key] {
							fmt.Println("found " + key + " for item " + item)
							shinyGoldParents++
							itemsToLookFor = append(itemsToLookFor, key)
							found[key] = true
						}
					}
				}
			}
			itemsToLookFor = itemsToLookFor[1:]
		}
		if len(itemsToLookFor) == 0 {
			break
		}
	}

	return shinyGoldParents
}

func partTwo() int {
	file, err := os.Open("test_input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	contains := make(map[string]map[string]int)
	totalBags := 0
	for scanner.Scan() {
		l := scanner.Text()
		subStrings := strings.Split(l, " bags contain ")
		if subStrings[1] == "no other bags." {
			continue
		}
		bagType := subStrings[0]
		contains[bagType] = make(map[string]int)
		contents := strings.Split(subStrings[1], ", ")

		for _, v := range contents {
			v = strings.TrimSuffix(v, " bags.")
			v = strings.TrimSuffix(v, " bag.")
			subParts := strings.SplitN(v, " ", 2)
			num, err := strconv.Atoi(subParts[0])
			content := subParts[1]
			if err != nil {
				fmt.Printf("Failed to parse %s\n", l)
				break
			}
			contains[bagType][content] = num
		}
	}
	newlyFoundBags := map[string]int{
		"shiny gold": 1,
	}
	allBagsFound := make(map[string]int)

	for {
		nextCycle := make(map[string]int)
		for k, v := range newlyFoundBags {
			children := contains[k]
			for c, count := range children {
				allBagsFound[c] += count * v
				nextCycle[c] += count * v
			}
		}
		if len(nextCycle) == 0 {
			break
		}
		newlyFoundBags = nextCycle
	}
	for _, v := range allBagsFound {
		totalBags += v
	}
	return totalBags
}

func main() {
	// fmt.Println(partOne())
	fmt.Println(partTwo())
}
