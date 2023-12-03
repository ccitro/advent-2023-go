package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var puzzle []string

func LoadPuzzle(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		puzzle = append(puzzle, line)
	}
}

func Part1() {
	sum := 0
	for _, v := range puzzle {
		lNum := -1
		rNum := -1
		for i := 0; i < len(v); i++ {
			if lNum == -1 {
				if thisNum, err := strconv.Atoi(string(v[i])); err == nil {
					lNum = thisNum
				}
			}

			if rNum == -1 {
				if thisNum, err := strconv.Atoi(string(v[len(v)-i-1])); err == nil {
					rNum = thisNum
				}
			}

			if lNum != -1 && rNum != -1 {
				break
			}
		}

		fmt.Printf("Line=%s, lNum=%d, rNum=%d\n", v, lNum, rNum)

		val := 10*lNum + rNum
		sum += val
	}

	println(sum)
}

func extractNum(s string) int {
	intVal, err := strconv.Atoi(string(s[0]))
	if err == nil {
		return intVal
	}

	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, v := range digits {
		if strings.HasPrefix(s, v) {
			return i + 1
		}
	}

	return -1
}

func Part2() {
	sum := 0

	for _, v := range puzzle {
		lNum := -1
		rNum := -1
		for i := 0; i < len(v); i++ {
			if lNum == -1 {
				thisNum := extractNum(v[i:])
				if thisNum != -1 {
					lNum = thisNum
				}
			}

			if rNum == -1 {
				thisNum := extractNum(v[len(v)-i-1:])
				if thisNum != -1 {
					rNum = thisNum
				}
			}

			if lNum != -1 && rNum != -1 {
				break
			}
		}

		fmt.Printf("Line=%s, lNum=%d, rNum=%d\n", v, lNum, rNum)

		val := 10*lNum + rNum
		sum += val
	}

	println(sum)
}
