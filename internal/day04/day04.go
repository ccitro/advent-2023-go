package day04

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	NumberSets [][]int
}

type Part2Card struct {
	*Card
	Matches int
	ID      int
}

var puzzle []Card

func LoadPuzzle(scanner *bufio.Scanner) {
	puzzle = make([]Card, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		suffix := strings.Split(line, ": ")[1]
		cardStrings := strings.Split(suffix, " | ")

		card := Card{}
		card.NumberSets = make([][]int, 0)
		for _, cardString := range cardStrings {
			numberList := strings.Split(cardString, " ")
			numberSet := make([]int, 0)
			for _, numberString := range numberList {
				if numberString == "" {
					continue
				}
				number, err := strconv.Atoi(numberString)
				if err != nil {
					panic(err)
				}
				numberSet = append(numberSet, number)
			}

			card.NumberSets = append(card.NumberSets, numberSet)
		}

		puzzle = append(puzzle, card)
	}
}

func getMatchCount(card *Card) int {
	winningSet := make(map[int]bool)
	for _, number := range card.NumberSets[0] {
		winningSet[number] = true
	}

	matches := 0
	for _, number := range card.NumberSets[1] {
		if _, ok := winningSet[number]; ok {
			matches++
		}
	}

	return matches
}

func scoreCard(card *Card) int {
	matches := getMatchCount(card)

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2.0, float64(matches-1)))
}

func Part1() {
	points := 0
	for _, card := range puzzle {
		points += scoreCard(&card)
	}

	println(points)
}

func getCardIDsWon(card *Part2Card) []int {
	// fmt.Printf("card #%d has %d matches\n", card.ID, card.Matches)
	cardIDs := make([]int, card.Matches)
	for i := range cardIDs {
		cardIDs[i] = card.ID + i + 1
	}

	return cardIDs
}

func Part2() {
	part2Puzzle := make([]Part2Card, len(puzzle)+1)
	for i, card := range puzzle {
		part2Puzzle[i+1] = Part2Card{&card, getMatchCount(&card), i + 1}
	}

	stack := make([]int, 0)
	for _, card := range part2Puzzle {
		if card.ID == 0 {
			continue
		}
		stack = append(stack, card.ID)
	}

	cardCount := 0
	for len(stack) > 0 {
		// fmt.Printf("stack: %v\n", stack)
		thisId := stack[0]
		stack = stack[1:]

		cardCount++

		thisCard := part2Puzzle[thisId]

		// fmt.Printf("stack before: %v\n", stack)
		stack = append(stack, getCardIDsWon(&thisCard)...)
		// fmt.Printf("stack after: %v\n", stack)
	}

	println(cardCount)
}
