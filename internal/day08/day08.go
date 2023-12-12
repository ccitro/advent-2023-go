package day08

import (
	"bufio"
	"fmt"
	"math"
)

const STEP_LEFT = 0
const STEP_RIGHT = 1

type Puzzle struct {
	Steps []int
	Nodes map[string][]string
}

var puzzle Puzzle

func LoadPuzzle(scanner *bufio.Scanner) {
	puzzle = Puzzle{
		Nodes: make(map[string][]string),
		Steps: make([]int, 0),
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if len(puzzle.Steps) == 0 {
			for _, c := range line {
				if c == 'L' {
					puzzle.Steps = append(puzzle.Steps, STEP_LEFT)
				} else {
					puzzle.Steps = append(puzzle.Steps, STEP_RIGHT)
				}
			}

			continue
		}

		label := line[0:3]
		left := line[7:10]
		right := line[12:15]

		puzzle.Nodes[label] = []string{left, right}
	}
}

func Part1() {
	steps := 0
	stepsLen := len(puzzle.Steps)
	currNode := "AAA"
	destNode := "ZZZ"

	for currNode != destNode {
		d := puzzle.Steps[steps%stepsLen]
		currNode = puzzle.Nodes[currNode][d]
		steps++
	}

	println(steps)
}

func Part2() {
	cycles := make([]int, 0)

	for n := range puzzle.Nodes {
		if n[2] != 'A' {
			continue
		}

		fmt.Printf("Checking %v\n", n)

		node := n
		i := 0
		for {
			n := i % len(puzzle.Steps)
			d := puzzle.Steps[n]
			node = puzzle.Nodes[node][d]
			if node[2] == 'Z' {
				cycles = append(cycles, i+1)
				break
			}

			i++
		}
	}

	println(lcm(cycles))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, int(math.Mod(float64(a), float64(b)))
	}

	return a
}

func lcm(nums []int) int {
	lcm := 1
	for _, n := range nums {
		lcm = lcm * n / gcd(lcm, n)
	}

	return lcm
}
