package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
}

type InstructionSet []Instruction

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne() int {
	file, err := os.Open("test_input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	instructions := make([]Instruction, 0)
	instructionHits := map[int]bool{}
	for scanner.Scan() {
		l := scanner.Text()
		operation := l[0:3]
		argument, err := strconv.Atoi(strings.TrimSpace(l[3:len(l)]))
		check(err)
		instruction := Instruction{
			operation: operation,
			argument:  argument,
		}
		instructions = append(instructions, instruction)
	}
	sequence := 0
	globalAcc := 0
	for {
		instruction := instructions[sequence]
		// here we need to check whether this instruction was already executed, if yes, we bail out
		_, ok := instructionHits[sequence]
		if ok {
			break
		}
		instructionHits[sequence] = true
		switch instruction.operation {
		case "acc":
			globalAcc += instruction.argument
			sequence += 1
		case "nop":
			sequence += 1
		case "jmp":
			// calculate new sequence
			sequence += instruction.argument
		}
	}
	return globalAcc
}

func endlessLoop(instructions []Instruction) (bool, int) {
	instructionHits := map[int]bool{}
	acc := 0
	sequence := 0
	for {
		if sequence >= len(instructions) {
			return false, acc
		}
		instruction := instructions[sequence]
		// here we need to check whether this instruction was already executed, if yes, we bail out
		_, ok := instructionHits[sequence]
		if ok {
			return true, 0
		}
		instructionHits[sequence] = true
		switch instruction.operation {
		case "acc":
			acc += instruction.argument
			sequence++
		case "nop":
			sequence++
		case "jmp":
			// calculate new sequence
			sequence += instruction.argument
		}
	}
}

func partTwo() int {
	file, err := os.Open("input.txt")
	defer file.Close()
	check(err)
	scanner := bufio.NewScanner(file)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		l := scanner.Text()
		operation := l[0:3]
		argument, err := strconv.Atoi(strings.TrimSpace(l[3:len(l)]))
		check(err)
		instruction := Instruction{
			operation: operation,
			argument:  argument,
		}
		instructions = append(instructions, instruction)
	}
	globalAcc := 0
	last := 0
	replaced := ""
	index := 0
	for {
		// Change next jmp or nop instruction
		if index != -1 {
			if replaced != "" {
				instructions[last].operation = replaced
			}
			if instructions[index].operation == "jmp" {
				instructions[index].operation = "nop"
				replaced = "jmp"
			} else {
				instructions[index].operation = "jmp"
				replaced = "nop"
			}
			last = index
		} else {
			log.Fatal("I couldn't find any more instructions to replace!")
		}
		endless, acc := endlessLoop(instructions)
		if !endless {
			globalAcc = acc
			break
		}
		index = findNext(instructions, last)
	}
	return globalAcc
}

func findNext(instructions []Instruction, pos int) int {
	for i, instruction := range instructions[pos+1:] {
		if instruction.operation == "jmp" || instruction.operation == "nop" {
			return i + pos + 1
		}
	}
	return -1
}

func main() {
	// fmt.Println(partOne())
	fmt.Println(partTwo())
}
