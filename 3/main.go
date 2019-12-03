package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

func main() {
	f, err := os.Open("input.txt")
	handleErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	firstOps := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	secondOps := strings.Split(scanner.Text(), ",")

	crossedFirst := getCrossed(firstOps)
	crossedSecond := getCrossed(secondOps)

	minDist := math.MaxInt64
	minSteps := math.MaxInt64
	for point := range crossedFirst {
		_, ok := crossedSecond[point]
		if ok {
			if distance(point) < minDist {
				minDist = distance(point)
			}
		}
	}
	fmt.Printf("Part 1: Minimum distance is %d\n", minDist)
	for point, stepsFirst := range crossedFirst {
		stepsSecond, ok := crossedSecond[point]
		if ok {
			if stepsFirst+stepsSecond < minSteps {
				minSteps = stepsFirst + stepsSecond
			}
		}
	}
	fmt.Printf("Part 2: Minimum steps is %d\n", minSteps)

}

func getCrossed(ops []string) map[point]int {
	crossed := make(map[point]int)
	var (
		state point
		steps int
	)
	for _, op := range ops {
		var mover func(point) point
		switch op[0] {
		case 'U':
			mover = func(p point) point {
				p.Y++
				return p
			}
		case 'D':
			mover = func(p point) point {
				p.Y--
				return p
			}
		case 'R':
			mover = func(p point) point {
				p.X++
				return p
			}
		case 'L':
			mover = func(p point) point {
				p.X--
				return p
			}
		}
		count, err := strconv.Atoi(op[1:])
		handleErr(err)
		for i := 0; i < count; i++ {
			state = mover(state)
			steps++
			crossed[state] = steps
		}
	}
	return crossed
}

func distance(a point) int {
	return abs(a.X) + abs(a.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func handleErr(err error) {
	if err != nil {
		log.Fatalf("Fatal due to err: %s", err.Error())
	}
}
