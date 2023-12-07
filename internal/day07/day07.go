package day07

import (
	"bufio"
	"math"
	"sort"
	"strconv"
	"strings"
)

type HandType int

var (
	HandTypeFiveofAKind  HandType = 1
	HandTypeFourOfAKind  HandType = 2
	HandTypeFullHouse    HandType = 3
	HandTypeThreeOfAKind HandType = 4
	HandTypeTwoPair      HandType = 5
	HandTypeOnePair      HandType = 6
	HandTypeHighCard     HandType = 7
)

func GetHandTypeLabel(handType HandType) string {
	switch handType {
	case HandTypeFiveofAKind:
		return "Five of a kind"
	case HandTypeFourOfAKind:
		return "Four of a kind"
	case HandTypeFullHouse:
		return "Full house"
	case HandTypeThreeOfAKind:
		return "Three of a kind"
	case HandTypeTwoPair:
		return "Two pair"
	case HandTypeOnePair:
		return "One pair"
	case HandTypeHighCard:
		return "High card"
	default:
		return "Unknown"
	}
}

type Hand struct {
	Cards    Cards
	HandType HandType
	Bid      int
}

type Cards = []rune
type Player struct {
	Cards Cards
	Bid   int
}
type Puzzle = []Player

var puzzle Puzzle

func LoadPuzzle(scanner *bufio.Scanner) {
	puzzle = make(Puzzle, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		player := Player{}
		parts := strings.Split(line, " ")
		cards := make(Cards, 0)
		for _, c := range parts[0] {
			cards = append(cards, c)
		}
		player.Cards = cards

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		player.Bid = bid

		puzzle = append(puzzle, player)
	}
}

func classifyHand(cards Cards) HandType {
	mappedCards := make(map[rune]int)
	for _, card := range cards {
		mappedCards[card]++
	}

	if mappedCards[cards[0]] == 5 {
		return HandTypeFiveofAKind
	}

	if len(mappedCards) == 2 {
		for _, count := range mappedCards {
			if count == 4 {
				return HandTypeFourOfAKind
			}

			if count == 3 {
				return HandTypeFullHouse
			}
		}
	}

	if len(mappedCards) == 3 {
		for _, count := range mappedCards {
			if count == 3 {
				return HandTypeThreeOfAKind
			}
		}
		return HandTypeTwoPair
	}

	if len(mappedCards) == 4 {
		return HandTypeOnePair
	}

	return HandTypeHighCard
}

func rankCard(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func (h *Hand) GreaterThan(other *Hand) bool {
	if h.HandType < other.HandType {
		return true
	}

	if h.HandType > other.HandType {
		return false
	}

	for i := 0; i < len(h.Cards); i++ {
		if rankCard(h.Cards[i]) > rankCard(other.Cards[i]) {
			return true
		}

		if rankCard(h.Cards[i]) < rankCard(other.Cards[i]) {
			return false
		}
	}

	return false
}

func Part1() {
	var hands []Hand
	for _, player := range puzzle {
		hand := Hand{
			Cards:    player.Cards,
			Bid:      player.Bid,
			HandType: classifyHand(player.Cards),
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].GreaterThan(&hands[j])
	})

	totalWinnings := 0
	for i, hand := range hands {
		rank := len(hands) - i
		score := rank * hand.Bid
		// fmt.Printf(
		// 	"Rank %d - cards: %s, type: %s, bid %d, score: %d\n",
		// 	rank, string(hand.Cards), GetHandTypeLabel(hand.HandType), hand.Bid, score,
		// )
		totalWinnings += score
	}

	println(totalWinnings)
}

func generateWildcardPermutations(cards Cards) []Cards {
	wildcardSlots := make([]int, 0)
	for i, card := range cards {
		if card == 'J' {
			wildcardSlots = append(wildcardSlots, i)
		}
	}

	if len(wildcardSlots) == 0 {
		return []Cards{cards}
	}

	uniqueCards := 12
	numPermutations := int(math.Pow(float64(uniqueCards), float64(len(wildcardSlots))))

	permutations := make([]Cards, numPermutations)
	for i := 0; i < int(numPermutations); i++ {
		// break down the permutation index into a base 12 number
		permutation := make([]int, len(wildcardSlots))
		permutationIndex := i
		for j := 0; j < len(wildcardSlots); j++ {
			permutation[j] = permutationIndex % uniqueCards
			permutationIndex /= uniqueCards
		}

		// convert the base 12 number into a permutation
		permutations[i] = make(Cards, len(cards))
		copy(permutations[i], cards)
		for j, slot := range wildcardSlots {
			var c rune
			// 0 = '2', 1 = '3', ... 9 = 'T', 10 = 'Q', 11 = 'K', 12 = 'A'
			if permutation[j] < 8 {
				c = rune(permutation[j] + '2')
			} else if permutation[j] == 8 {
				c = 'T'
			} else if permutation[j] == 9 {
				c = 'Q'
			} else if permutation[j] == 10 {
				c = 'K'
			} else if permutation[j] == 11 {
				c = 'A'
			} else {
				panic("invalid permutation")
			}

			permutations[i][slot] = c
		}
	}

	return permutations
}

func part2ClassifyHand(cards Cards) HandType {
	permutations := generateWildcardPermutations(cards)

	bestHandType := HandTypeHighCard
	for _, permutation := range permutations {
		handType := classifyHand(permutation)
		if handType < bestHandType {
			bestHandType = handType
		}
	}

	return bestHandType
}

func part2RankCard(card rune) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}

func (h *Hand) Part2GreaterThan(other *Hand) bool {
	if h.HandType < other.HandType {
		return true
	}

	if h.HandType > other.HandType {
		return false
	}

	for i := 0; i < len(h.Cards); i++ {
		if part2RankCard(h.Cards[i]) > part2RankCard(other.Cards[i]) {
			return true
		}

		if part2RankCard(h.Cards[i]) < part2RankCard(other.Cards[i]) {
			return false
		}
	}

	return false
}

func Part2() {
	var hands []Hand
	for _, player := range puzzle {
		hand := Hand{
			Cards:    player.Cards,
			Bid:      player.Bid,
			HandType: part2ClassifyHand(player.Cards),
		}
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Part2GreaterThan(&hands[j])
	})

	totalWinnings := 0
	for i, hand := range hands {
		rank := len(hands) - i
		score := rank * hand.Bid
		// fmt.Printf(
		// 	"Rank %d - cards: %s, type: %s, bid %d, score: %d\n",
		// 	rank, string(hand.Cards), GetHandTypeLabel(hand.HandType), hand.Bid, score,
		// )
		totalWinnings += score
	}

	println(totalWinnings)
}
