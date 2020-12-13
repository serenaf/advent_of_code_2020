package main

import (
	"bufio"
	"fmt"
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
	file, err := os.Open("input.txt")
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

func main() {
	fmt.Println(partOne())
}
