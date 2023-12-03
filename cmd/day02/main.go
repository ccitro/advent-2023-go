package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Day02Pull = map[string]int
type Day02Game struct {
	Game  int
	Pulls []Day02Pull
}

var puzzle []Day02Game

func loadPuzzle(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ": ")

		gameNum, err := strconv.Atoi(parts[0][5:])
		if err != nil {
			panic(err)
		}

		handfulls := strings.Split(parts[1], "; ")
		pulls := make([]Day02Pull, len(handfulls))
		for i, v := range handfulls {
			pulls[i] = make(Day02Pull)
			clusters := strings.Split(v, ", ")
			for _, cluster := range clusters {
				parts := strings.Split(cluster, " ")
				n, err := strconv.Atoi(parts[0])
				if err != nil {
					panic(err)
				}

				pulls[i][parts[1]] = n
			}
		}

		puzzle = append(puzzle, Day02Game{Game: gameNum, Pulls: pulls})
	}
}

func part1() {
	var maxPullsByLabel = make(map[string]int)
	maxPullsByLabel["red"] = 12
	maxPullsByLabel["green"] = 13
	maxPullsByLabel["blue"] = 14

	sumOfGoodGames := 0
	for _, game := range puzzle {
		badGame := false
		for _, pull := range game.Pulls {
			for label, num := range pull {
				if num > maxPullsByLabel[label] {
					badGame = true
					break
				}
			}
		}

		if !badGame {
			sumOfGoodGames += game.Game
		}
	}

	println(sumOfGoodGames)
}

func part2() {
	powerSum := 0
	for _, game := range puzzle {
		var largestPullsByLabel = make(map[string]int)
		for _, pull := range game.Pulls {
			for label, num := range pull {
				if num > largestPullsByLabel[label] {
					largestPullsByLabel[label] = num
				}
			}
		}
		gamePower := 1
		for _, v := range largestPullsByLabel {
			gamePower *= v
		}
		powerSum += gamePower
	}

	println(powerSum)
}

func main() {
	filename := "input.txt"
	method := part1
	for _, v := range os.Args {
		if v == "part2" || v == "2" {
			method = part2
		}
		if strings.HasSuffix(v, ".txt") {
			filename = v
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	loadPuzzle(file)
	method()
}
