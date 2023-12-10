package day

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "07", "input"))
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}

	if err := part1(string(in)); err != nil {
		return fmt.Errorf("failed to run part1: %w", err)
	}
	if err := part2(string(in)); err != nil {
		return fmt.Errorf("failed to run part2: %w", err)
	}

	return nil
}

var part1cardMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var part2cardMap = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

type handBid struct {
	cardSort string
	hand     string
	bid      int
}

func part1(in string) error {
	rankings := make([][]handBid, 7)
	for _, l := range strings.Split(in, "\n") {
		split := strings.Split(l, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])

		cards := make([]int, 15)
		cardSort := ""
		for _, card := range hand {
			cardVal := part1cardMap[card]
			cards[cardVal] += 1
			cardSort += string(rune(cardVal))
		}

		sort.Slice(cards, func(i, j int) bool {
			return cards[i] > cards[j]
		})

		hb := handBid{
			cardSort: cardSort,
			bid:      bid,
		}
		switch cards[0] {
		case 1:
			rankings[0] = append(rankings[0], hb)
		case 2:
			if cards[1] == 2 {
				rankings[2] = append(rankings[2], hb)
			} else {
				rankings[1] = append(rankings[1], hb)
			}
		case 3:
			if cards[1] == 2 {
				rankings[4] = append(rankings[4], hb)
			} else {
				rankings[3] = append(rankings[3], hb)
			}
		case 4:
			rankings[5] = append(rankings[5], hb)
		case 5:
			rankings[6] = append(rankings[6], hb)
		}
	}

	for _, ranking := range rankings {
		sort.Slice(ranking, func(i, j int) bool {
			return strings.Compare(ranking[i].cardSort, ranking[j].cardSort) < 0
		})
	}

	acc := 0
	rank := 1
	for _, ranking := range rankings {
		for _, r := range ranking {
			acc += r.bid * rank
			rank += 1
		}
	}

	fmt.Printf("Part One: %d\n", acc)

	return nil
}

func part2(in string) error {
	rankings := make([][]handBid, 7)
	for _, l := range strings.Split(in, "\n") {
		split := strings.Split(l, " ")
		hand := split[0]
		bid, _ := strconv.Atoi(split[1])

		cards := make([]int, 15)
		cardSort := ""
		numJs := 0
		for _, card := range hand {
			cardVal := part2cardMap[card]
			if card != 'J' {
				cards[cardVal] += 1
			} else {
				numJs++
			}
			cardSort += string(rune(cardVal))
		}

		sort.Slice(cards, func(i, j int) bool {
			return cards[i] > cards[j]
		})

		hb := handBid{
			cardSort: cardSort,
			bid:      bid,
		}
		switch cards[0] + numJs {
		case 1:
			rankings[0] = append(rankings[0], hb)
		case 2:
			if cards[1] == 2 {
				rankings[2] = append(rankings[2], hb)
			} else {
				rankings[1] = append(rankings[1], hb)
			}
		case 3:
			if cards[1] == 2 {
				rankings[4] = append(rankings[4], hb)
			} else {
				rankings[3] = append(rankings[3], hb)
			}
		case 4:
			rankings[5] = append(rankings[5], hb)
		case 5:
			rankings[6] = append(rankings[6], hb)
		}
	}

	for _, ranking := range rankings {
		sort.Slice(ranking, func(i, j int) bool {
			return strings.Compare(ranking[i].cardSort, ranking[j].cardSort) < 0
		})
	}

	acc := 0
	rank := 1
	for _, ranking := range rankings {
		for _, r := range ranking {
			acc += r.bid * rank
			rank += 1
		}
	}

	fmt.Printf("Part Two: %d\n", acc)

	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "7",
			RunE: day,
		},
		"2023",
		"7",
	)
}
