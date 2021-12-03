package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day1/day1.in")
	scanner := bufio.NewScanner(file)
	measurements := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		m, _ := strconv.Atoi(line)
		measurements = append(measurements, m)
	}

	previous := 0
	increases := -1
	for i := 0; i < len(measurements)-2; i++ {
		windowSum := measurements[i] + measurements[i+1] + measurements[i+2]
		if windowSum > previous {
			increases++
		}
		previous = windowSum
	}
	fmt.Println(increases)
}
