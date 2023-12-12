package main

import (
	"github.com/ccitro/advent-2023-go/internal/day01"
	"github.com/ccitro/advent-2023-go/internal/day02"
	"github.com/ccitro/advent-2023-go/internal/day03"
	"github.com/ccitro/advent-2023-go/internal/day04"
	"github.com/ccitro/advent-2023-go/internal/day05"
	"github.com/ccitro/advent-2023-go/internal/day06"
	"github.com/ccitro/advent-2023-go/internal/day07"
	"github.com/ccitro/advent-2023-go/internal/day08"

	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Loader = func(scanner *bufio.Scanner)
type DayMethods struct {
	LoadPuzzle Loader
	Part1      func()
	Part2      func()
}

var dayMethods = map[string]DayMethods{
	"day08": {LoadPuzzle: day08.LoadPuzzle, Part1: day08.Part1, Part2: day08.Part2},
	"day07": {LoadPuzzle: day07.LoadPuzzle, Part1: day07.Part1, Part2: day07.Part2},
	"day06": {LoadPuzzle: day06.LoadPuzzle, Part1: day06.Part1, Part2: day06.Part2},
	"day05": {LoadPuzzle: day05.LoadPuzzle, Part1: day05.Part1, Part2: day05.Part2},
	"day04": {LoadPuzzle: day04.LoadPuzzle, Part1: day04.Part1, Part2: day04.Part2},
	"day03": {LoadPuzzle: day03.LoadPuzzle, Part1: day03.Part1, Part2: day03.Part2},
	"day02": {LoadPuzzle: day02.LoadPuzzle, Part1: day02.Part1, Part2: day02.Part2},
	"day01": {LoadPuzzle: day01.LoadPuzzle, Part1: day01.Part1, Part2: day01.Part2},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <day> [part2] [input.txt]")
		return
	}

	// allowed formats: day01, 1, day1
	dayArg := os.Args[1]

	re := regexp.MustCompile(`\d+`)
	day := re.FindString(dayArg)
	if day == "" {
		fmt.Printf("Unknown day: %s\n", dayArg)
		return
	}

	if len(day) == 1 {
		day = fmt.Sprintf("0%s", day)
	}

	day = fmt.Sprintf("day%s", day)

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
