package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var moduleTotal, fuelTotal int
	f, err := os.Open("input.txt")
	handleErr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		handleErr(err)
		moduleFuel := calcFuel(mass)
		moduleTotal += moduleFuel
		for additionalFuel := moduleFuel;;{
			additionalFuel = calcFuel(additionalFuel)
			if additionalFuel < 0 {
				break
			}
			moduleFuel+= additionalFuel
		}
		fuelTotal += moduleFuel
	}
	fmt.Printf("Part 1: Total fuel for just the modules is %d!\n", moduleTotal)
	fmt.Printf("Part 2: Total fuel for also flying out the fuels is %d!\n", fuelTotal)
}

func calcFuel(mass int) int {
	return mass/3-2
}

func handleErr(err error) {
	if err != nil {
		log.Fatalf("Fatal due to err: %s", err.Error())
	}
}