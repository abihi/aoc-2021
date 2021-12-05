package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("day5/day5.in")
	scanner := bufio.NewScanner(file)

	pointMap := map[Point]int{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "->")
		x1y1 := strings.Split(strings.TrimSpace(split[0]), ",")
		x2y2 := strings.Split(strings.TrimSpace(split[1]), ",")
		point1 := Point{x: strToInt(x1y1[0]), y: strToInt(x1y1[1])}
		point2 := Point{x: strToInt(x2y2[0]), y: strToInt(x2y2[1])}
		pointMap = fillWithLines(pointMap, point1, point2)
	}

	overlaps := 0
	for _, v := range pointMap {
		if v > 1 {
			overlaps++
		}
	}

	fmt.Println("ans:", overlaps)
}

func fillWithDiagonal(pointMap map[Point]int, p1 Point, p2 Point) map[Point]int {
	if p1.x > p2.x {
		p2, p1 = p1, p2
	}

	slope := (p2.y - p1.y) / (p2.x - p1.x)
	j := p1.y
	for i := p1.x + 1; i < p2.x; i++ {
		j += slope
		point := Point{i, j}
		pointMap[point]++
	}

	return pointMap
}

func fillWithLines(pointMap map[Point]int, p1 Point, p2 Point) map[Point]int {
	pointMap[p1]++
	pointMap[p2]++

	if p1.x != p2.x && p1.y != p2.y {
		return fillWithDiagonal(pointMap, p1, p2)
	}

	if p1.x > p2.x {
		p2, p1 = p1, p2
	}

	var point Point
	for i := p1.x + 1; i < p2.x; i++ {
		point = Point{i, p1.y}
		pointMap[point]++
	}

	if p1.y > p2.y {
		p2, p1 = p1, p2
	}
	for i := p1.y + 1; i < p2.y; i++ {
		point = Point{p2.x, i}
		pointMap[point]++
	}

	return pointMap
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
