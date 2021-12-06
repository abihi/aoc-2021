package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day6/day6.in")
	scanner := bufio.NewScanner(file)

	simulation := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, n := range strings.Split(line, ",") {
			simulation[strToInt(n)] += 1
		}
	}

	fmt.Println("Initial state:", simulation)
	for day := 0; day < 256; day++ {
		newSimulation := map[int]int{}
		for k, v := range simulation {
			if k == 0 {
				newSimulation[6] += v
				newSimulation[8] = v
			} else {
				newSimulation[k-1] += v
			}
		}
		simulation = newSimulation
		//fmt.Printf("After %d days: %d\n", day, simulation)
	}

	fishes := 0
	for _, v := range simulation {
		fishes += v
	}
	fmt.Println("Ans: ", fishes)
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
