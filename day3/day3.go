package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("day3/day3.in")
	scanner := bufio.NewScanner(file)
	oxygenRating := []string{}
	co2Rating := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		oxygenRating = append(oxygenRating, line)
		co2Rating = append(co2Rating, line)
	}

	oxygenRating = calculateRating(oxygenRating, '1', '0')
	co2Rating = calculateRating(co2Rating, '0', '1')

	powerConsumption := binaryStringToInt(oxygenRating[0]) * binaryStringToInt(co2Rating[0])
	fmt.Println(powerConsumption)
}

func calculateRating(ratings []string, keepChar1 byte, keepChar2 byte) []string {
	i := 0
	for len(ratings) > 1 {
		ones, zeros := count1sAnd0s(ratings, i)
		if ones >= zeros {
			ratings = keep(i, keepChar1, ratings)
		} else {
			ratings = keep(i, keepChar2, ratings)
		}
		i++
	}
	return ratings
}

func count1sAnd0s(ratings []string, pos int) (int, int) {
	ones := 0
	zeros := 0
	for _, rating := range ratings {
		if rating[pos] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return ones, zeros
}

func keep(pos int, char byte, ratings []string) []string {
	if len(ratings) == 1 {
		return ratings
	}
	newRating := []string{}
	for _, rating := range ratings {
		if rating[pos] == char {
			newRating = append(newRating, rating)
		}
	}
	return newRating
}

func binaryStringToInt(str string) int64 {
	i, err := strconv.ParseInt(str, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}
