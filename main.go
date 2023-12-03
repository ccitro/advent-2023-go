package main

import (
	"github.com/ccitro/advent-2023-go/internal/day01"
	"github.com/ccitro/advent-2023-go/internal/day02"
	"github.com/ccitro/advent-2023-go/internal/day03"

	"bufio"
	"fmt"
	"os"
	"strings"
)

type Loader = func(scanner *bufio.Scanner)
type DayMethods struct {
	LoadPuzzle Loader
	Part1      func()
	Part2      func()
}

var dayMethods = map[string]DayMethods{
	"day03": {LoadPuzzle: day03.LoadPuzzle, Part1: day03.Part1, Part2: day03.Part2},
	"day02": {LoadPuzzle: day02.LoadPuzzle, Part1: day02.Part1, Part2: day02.Part2},
	"day01": {LoadPuzzle: day01.LoadPuzzle, Part1: day01.Part1, Part2: day01.Part2},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day> [part2] [input.txt]")
		return
	}

	day := os.Args[1]
	if dayMethods[day].LoadPuzzle == nil {
		fmt.Printf("Unknown day: %s\n", day)
		return
	}

	launch(day, dayMethods[day])
}

func launch(name string, dayMethods DayMethods) {
	filename := "example.txt"
	method := dayMethods.Part1
	for _, v := range os.Args {
		if v == "part2" || v == "2" {
			method = dayMethods.Part2
		}
		if strings.HasSuffix(v, ".txt") {
			filename = v
		}
	}

	filePath := fmt.Sprintf("./internal/%s/%s", name, filename)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	dayMethods.LoadPuzzle(scanner)
	method()
}
