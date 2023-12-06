package day05

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct {
	Seeds []int
	Maps  []CategoryMap
}

type CategoryMap struct {
	From    string
	To      string
	Entries [][]int
}

var puzzle Puzzle

func LoadPuzzle(scanner *bufio.Scanner) {
	puzzle = Puzzle{}
	currentMap := CategoryMap{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if currentMap.From != "" {
				puzzle.Maps = append(puzzle.Maps, currentMap)
				currentMap = CategoryMap{}
			}
			continue
		}

		if puzzle.Seeds == nil {
			puzzle.Seeds = make([]int, 0)
			// seeds: 79 14 55 13
			l := line[7:]
			ids := strings.Split(l, " ")
			for _, id := range ids {
				n, err := strconv.Atoi(id)
				if err != nil {
					panic(err)
				}

				puzzle.Seeds = append(puzzle.Seeds, n)
			}
			continue
		}

		if strings.HasSuffix(line, " map:") {
			// soil-to-fertilizer map:
			parts := strings.Split(line, " ")
			labels := strings.Split(parts[0], "-")
			currentMap.From = labels[0]
			currentMap.To = labels[2]
			currentMap.Entries = make([][]int, 0)
			continue
		}

		// 0 15 37
		ids := strings.Split(line, " ")
		entry := make([]int, 0)
		for _, id := range ids {
			n, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}

			entry = append(entry, n)
		}

		currentMap.Entries = append(currentMap.Entries, entry)
	}

	if currentMap.From != "" {
		puzzle.Maps = append(puzzle.Maps, currentMap)
	}
}

func getCorrespondingId(id int, from string, to string) int {
	currentMap := CategoryMap{}
	for _, m := range puzzle.Maps {
		if m.From == from {
			currentMap = m
			break
		}
	}

	if currentMap.From == "" {
		panic(fmt.Sprintf("No map found for %s", from))
	}

	mappedId := getMappedId(id, &currentMap)

	if currentMap.To == to {
		return mappedId
	}

	return getCorrespondingId(mappedId, currentMap.To, to)
}

func getMappedId(id int, currentMap *CategoryMap) int {
	for _, entry := range currentMap.Entries {
		destStart := entry[0]
		sourceStart := entry[1]
		length := entry[2]

		if sourceStart <= id && id < sourceStart+length {
			return destStart + (id - sourceStart)
		}
	}

	// panic(fmt.Sprintf("No mapping found for %d from %s to %s", id, currentMap.From, currentMap.To))
	return id
}

func Part1() {
	min := 99999999999
	for _, seed := range puzzle.Seeds {
		location := getCorrespondingId(seed, "seed", "location")
		if location < min {
			min = location
		}
	}

	println(min)
}

func checkSeed(seedStart int, seedLen int, done chan int) {
	fmt.Printf("Checking seed %d %d\n", seedStart, seedLen)
	min := 99999999999
	for i := seedStart; i < seedStart+seedLen; i++ {
		location := getCorrespondingId(i, "seed", "location")
		if location < min {
			min = location
		}
	}

	done <- min
}

func Part2() {
	min := 99999999999
	messages := make(chan int, len(puzzle.Seeds)/2)

	for seedSeq := 0; seedSeq < len(puzzle.Seeds); seedSeq += 2 {
		go checkSeed(puzzle.Seeds[seedSeq], puzzle.Seeds[seedSeq+1], messages)
	}

	for i := 0; i < len(puzzle.Seeds)/2; i++ {
		thisMin := <-messages
		if thisMin < min {
			min = thisMin
		}
	}

	println(min)
}
