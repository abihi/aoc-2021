package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day7/day.in")
	scanner := bufio.NewScanner(file)

	positions := []int{}
	freq := map[int]int{}
	max := 0
	min := int(^uint(0) >> 1)

	scanner.Scan()
	line := strings.Split(scanner.Text(), ",")
	for _, num := range line {
		n := strToInt(num)
		positions = append(positions, n)
		freq[n]++
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}

	best := int(^uint(0) >> 1)
	bestNum := 0
	for i := min; i <= max; i++ {
		// mostFreq := mostFrequentNum(freq)
		totalFuel := calculateTotalFuel(positions, i)
		if totalFuel < best {
			best = totalFuel
			bestNum = i
		}
	}
	println(best, bestNum)
}

func calculateTotalFuel(positions []int, mostFreq int) int {
	totalFuel := 0
	for _, n := range positions {
		steps := abs(n - mostFreq)
		for i := 1; i <= steps; i++ {
			totalFuel += i
		}
	}
	return totalFuel
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mostFrequentNum(freq map[int]int) int {
	mostFreq := []int{0, 0}
	for k, v := range freq {
		if v > mostFreq[1] {
			mostFreq[0] = k
			mostFreq[1] = v
		}
	}
	delete(freq, mostFreq[0])
	return mostFreq[0]
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
