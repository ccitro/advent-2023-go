package day02

import (
	"bufio"
	"strconv"
	"strings"
)

type Pull = map[string]int
type Game struct {
	Game  int
	Pulls []Pull
}

var puzzle []Game

func LoadPuzzle(scanner *bufio.Scanner) {
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
		pulls := make([]Pull, len(handfulls))
		for i, v := range handfulls {
			pulls[i] = make(Pull)
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

		puzzle = append(puzzle, Game{Game: gameNum, Pulls: pulls})
	}
}

func Part1() {
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

func Part2() {
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
