package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("day2/day2.in")
	scanner := bufio.NewScanner(file)
	position := []int{0, 0, 0}
	for scanner.Scan() {
		line := scanner.Text()
		direction := strings.Split(line, " ")
		switch direction[0] {
		case "forward":
			position[0] += strToInt(direction[1])
			position[1] += position[2] * strToInt(direction[1])
		case "down":
			position[2] += strToInt(direction[1])
		case "up":
			position[2] -= strToInt(direction[1])
		}
	}

	fmt.Println(position[0] * position[1])
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
