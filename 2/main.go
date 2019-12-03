package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const target = 19690720

func main() {
	in, err := ioutil.ReadFile("input.txt")
	handleErr(err)
	ss := strings.Split(string(in), ",")
	var program []int
	for _, s := range ss {
		i, err := strconv.Atoi(s)
		handleErr(err)
		program = append(program, i)
	}
	fmt.Printf("Part 1: First opcode for 1202 is %d\n",
		gravity(12, 2, append([]int{}, program...)))
	var noun, verb int
NounVerbLoop:
	for noun = 0; noun < 99; noun++ {
		for verb = 0; verb < 99; verb++ {
			if gravity(noun, verb, append([]int{}, program...)) == target {
				break NounVerbLoop
			}
		}
	}
	fmt.Printf("Part 2: Target noun is %d target verb is %d and their sum is %d\n",
		noun, verb, 100*noun+verb)
}

func gravity(noun, verb int, program []int) int {
	program[1] = noun
	program[2] = verb
	for i := 0; program[i] != 99; i += 4 {
		switch program[i] {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		}
	}
	return program[0]
}
func handleErr(err error) {
	if err != nil {
		log.Fatalf("Fatal due to err: %s", err.Error())
	}
}
