package day03

import (
	"bufio"
	"fmt"
	"strconv"
)

var puzzle [][]rune

func LoadPuzzle(scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		runes := []rune(line)
		puzzle = append(puzzle, runes)
	}
}

func isInBounds(x, y int) bool {
	if y < 0 || y >= len(puzzle) {
		return false
	}

	if x < 0 || x >= len(puzzle[y]) {
		return false
	}

	return true
}

func isDigit(x, y int) bool {
	if !isInBounds(x, y) {
		return false
	}

	return puzzle[y][x] >= '0' && puzzle[y][x] <= '9'
}

func isGear(x, y int) bool {
	if !isInBounds(x, y) {
		return false
	}

	return puzzle[y][x] == '*'
}

func isSymbol(x, y int) bool {
	if !isInBounds(x, y) {
		return false
	}

	if isDigit(x, y) {
		return false
	}

	return puzzle[y][x] != '.'
}

type PartNumber struct {
	startX, startY int
	endX, endY     int
	n              int
}

func (n *PartNumber) IsAdjacentToSymbol() bool {
	tlx := n.startX - 1
	tly := n.startY - 1
	brx := n.endX + 1
	bry := n.endY + 1

	for i := tlx; i <= brx; i++ {
		for j := tly; j <= bry; j++ {
			if isSymbol(i, j) {
				return true
			}
		}
	}

	return false
}

type Point struct {
	x, y int
}

func (n *PartNumber) GetAdjacentGears() []Point {
	points := make([]Point, 0)
	tlx := n.startX - 1
	tly := n.startY - 1
	brx := n.endX + 1
	bry := n.endY + 1

	for i := tlx; i <= brx; i++ {
		for j := tly; j <= bry; j++ {
			if isGear(i, j) {
				points = append(points, Point{x: i, y: j})
			}
		}
	}

	return points
}

func getPartNumber(x, y int) *PartNumber {
	digits := make([]rune, 0)
	for {
		if !isDigit(x, y) {
			break
		}

		digits = append(digits, puzzle[y][x])
		x++
	}

	if len(digits) == 0 {
		return nil
	}

	n, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(err)
	}

	return &PartNumber{
		startX: x - len(digits),
		startY: y,
		endX:   x - 1,
		endY:   y,
		n:      n,
	}
}

func Part1() {
	sum := 0
	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[y]); x++ {
			n := getPartNumber(x, y)
			if n == nil {
				continue
			}

			fmt.Printf("Found part number %d at (%d, %d)\n", n.n, n.startX, n.startY)
			if n.IsAdjacentToSymbol() {
				fmt.Printf("Part number %d is adjacent to a symbol\n", n.n)
				sum += n.n
			}

			x = n.endX
		}
	}
	println(sum)
}

func Part2() {
	gearPartNumbers := make(map[int][]PartNumber)
	for y := 0; y < len(puzzle); y++ {
		for x := 0; x < len(puzzle[y]); x++ {
			n := getPartNumber(x, y)
			if n == nil {
				continue
			}

			fmt.Printf("Found part number %d at (%d, %d)\n", n.n, n.startX, n.startY)
			adjacentGears := n.GetAdjacentGears()
			if len(adjacentGears) > 0 {
				fmt.Printf("Part number %d is adjacent to %d gears\n", n.n, len(adjacentGears))
				for _, gear := range adjacentGears {
					flatId := gear.y*len(puzzle[0]) + gear.x
					if _, ok := gearPartNumbers[flatId]; !ok {
						gearPartNumbers[flatId] = make([]PartNumber, 0)
					}

					gearPartNumbers[flatId] = append(gearPartNumbers[flatId], *n)
				}
			}
			x = n.endX
		}
	}

	ratioSum := 0
	for k, v := range gearPartNumbers {
		fmt.Printf("Gear at %d,%d has %d part numbers\n", k%len(puzzle[0]), k/len(puzzle[0]), len(v))
		if len(v) == 1 {
			fmt.Printf("Gear at %d,%d has only one part number\n", k%len(puzzle[0]), k/len(puzzle[0]))
			continue
		}

		ratio := 1
		for _, n := range v {
			ratio *= n.n
		}

		ratioSum += ratio
	}

	println(ratioSum)
}
