package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	rows []map[string]bool
	cols []map[string]bool
}

func main() {
	file, _ := os.Open("day4/day4.in")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	bingoNumbers := strings.Split(scanner.Text(), ",")
	boards := []Board{}

	matrix := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(matrix) > 0 {
				boards = append(boards, fillBoard(matrix))
			}
			matrix = [][]string{}
			continue
		}
		row := trim(strings.Split(line, " "))
		matrix = append(matrix, row)
	}

	i, lastNumber, _ := findBingoBoard(bingoNumbers, boards)
	fmt.Println("Part 1 answer: ", sumBoard(boards[i], lastNumber))

	lastBoard, remainingNumbers := findLastBingoBoard(boards, bingoNumbers)
	i, lastNumber, _ = findBingoBoard(remainingNumbers, lastBoard)
	fmt.Println("Part 2 answer: ", sumBoard(lastBoard[i], lastNumber))
}

func findLastBingoBoard(boards []Board, bingoNumbers []string) ([]Board, []string) {
	i := 0
	currPos := 0
	for len(boards) != 1 {
		i, _, currPos = findBingoBoard(bingoNumbers[currPos:], boards)
		boards = remove(boards, i)
	}
	return boards, bingoNumbers[currPos:]
}

func findBingoBoard(bingoNumbers []string, boards []Board) (int, int, int) {
	for i, number := range bingoNumbers {
		for j, board := range boards {
			for _, row := range board.rows {
				if _, ok := row[number]; ok {
					row[number] = true
				}
				if bingo(row) {
					return j, strToInt(number), i
				}
			}
			for _, col := range board.cols {
				if _, ok := col[number]; ok {
					col[number] = true
				}
				if bingo(col) {
					return j, strToInt(number), i
				}
			}
		}
	}
	panic("No bingo found")
}

func sumBoard(board Board, lastNumber int) int {
	sum := 0
	for _, row := range board.rows {
		for k, v := range row {
			if !v {
				sum += strToInt(k)
			}
		}
	}
	return sum * lastNumber
}

func bingo(vector map[string]bool) bool {
	matches := 0
	for _, v := range vector {
		if v {
			matches++
		}
	}
	if matches == len(vector) {
		return true
	}
	return false
}

func fillBoard(matrix [][]string) Board {
	board := Board{rows: []map[string]bool{}, cols: []map[string]bool{}}
	for i := 0; i < 5; i++ {
		row := map[string]bool{}
		col := map[string]bool{}
		for j := 0; j < 5; j++ {
			row[matrix[i][j]] = false
			col[matrix[j][i]] = false
		}
		board.rows = append(board.rows, row)
		board.cols = append(board.cols, col)
	}

	return board
}

func trim(str []string) []string {
	trimmed := []string{}
	for _, s := range str {
		if s == "" {
			continue
		}
		trimmed = append(trimmed, s)
	}
	return trimmed
}

func strToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func remove(b []Board, i int) []Board {
	b[i] = b[len(b)-1]
	return b[:len(b)-1]
}
