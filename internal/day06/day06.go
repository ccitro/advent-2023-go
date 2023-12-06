package day06

import (
	"bufio"
	"strconv"
	"strings"
)

type Puzzle struct {
	Times     []int
	Distances []int
}

var puzzle Puzzle

func LoadPuzzle(scanner *bufio.Scanner) {
	puzzle = Puzzle{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		intStrs := strings.Split(parts[1], " ")
		ints := make([]int, 0)
		for _, intStr := range intStrs {
			if intStr == "" {
				continue
			}
			n, err := strconv.Atoi(strings.Trim(intStr, " "))
			if err != nil {
				panic(err)
			}

			ints = append(ints, n)
		}

		if parts[0] == "Time" {
			puzzle.Times = ints
		} else {
			puzzle.Distances = ints
		}
	}
}

func calculateDistanceTraveled(time, holdTime int) int {
	travelTime := time - holdTime
	return travelTime * holdTime
}

func calculateWaysToWin(time, distance int) int {
	waysToWin := 0
	for holdTime := 1; holdTime < time-1; holdTime++ {
		thisDistance := calculateDistanceTraveled(time, holdTime)
		if thisDistance > distance {
			waysToWin++
		}
	}

	return waysToWin
}

func Part1() {
	margin := 1
	for i := 0; i < len(puzzle.Times); i++ {
		margin *= calculateWaysToWin(puzzle.Times[i], puzzle.Distances[i])
	}

	println(margin)
}

func Part2() {
	combinedTime := ""
	combinedDistance := ""

	for i := 0; i < len(puzzle.Times); i++ {
		combinedTime += strconv.Itoa(puzzle.Times[i])
		combinedDistance += strconv.Itoa(puzzle.Distances[i])
	}

	time, timeErr := strconv.Atoi(combinedTime)
	if timeErr != nil {
		panic(timeErr)
	}

	distance, distanceErr := strconv.Atoi(combinedDistance)
	if distanceErr != nil {
		panic(distanceErr)
	}

	println(calculateWaysToWin(time, distance))
}
