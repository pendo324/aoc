package day

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

type Color = string

const (
	RED   Color = "red"
	BLUE  Color = "blue"
	GREEN Color = "green"
)

type Game struct {
	Id   int
	Bags []Bag
}

type Bag = []Cube

type Cube struct {
	Color  Color
	Amount int
}

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "02", "input"))
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}

	games, err := parseData(string(in))
	if err != nil {
		return fmt.Errorf("failed to parse data: %w", err)
	}

	if err := part1(games); err != nil {
		return fmt.Errorf("failed to run part1: %w", err)
	}
	if err := part2(games); err != nil {
		return fmt.Errorf("failed to run part2: %w", err)
	}

	return nil
}

func parseData(in string) ([]Game, error) {
	games := []Game{}

	for _, l := range strings.Split(in, "\n") {
		s := strings.Split(l, ": ")
		id := strings.Split(s[0], "Game ")[1]

		idi, err := strconv.Atoi(id)
		if err != nil {
			return nil, fmt.Errorf("failed to convert to int: %w", err)
		}

		g := Game{
			Id: idi,
		}

		bags := []Bag{}
		bs := strings.Split(s[1], "; ")
		for _, b := range bs {
			bag := Bag{}
			cube := Cube{}

			ss := strings.Split(b, " ")
			for k := 0; k < len(ss); k = k + 2 {
				n := ss[k]
				c, _ := strings.CutSuffix(ss[k+1], ",")
				ni, err := strconv.Atoi(n)
				if err != nil {
					return nil, fmt.Errorf("failed to convert to int: %w", err)
				}

				cube.Amount = ni
				cube.Color = Color(c)
				bag = append(bag, cube)
			}

			bags = append(bags, bag)
		}

		g.Bags = bags

		games = append(games, g)
	}

	return games, nil
}

func part1(games []Game) error {
	possibleGames := []Game{}
	for _, g := range games {
		possible := true
		for _, b := range g.Bags {
			if !possible {
				break
			}
			for _, c := range b {
				if c.Color == RED && c.Amount > 12 ||
					c.Color == GREEN && c.Amount > 13 ||
					c.Color == BLUE && c.Amount > 14 {
					possible = false
					break
				}
			}
		}

		if possible {
			possibleGames = append(possibleGames, g)
		}
	}

	out := 0
	for _, g := range possibleGames {
		out += g.Id
	}

	fmt.Printf("Part One: %d\n", out)

	return nil
}

func part2(games []Game) error {
	acc := 0
	for _, g := range games {
		min := map[Color]int{
			RED:   0,
			GREEN: 0,
			BLUE:  0,
		}
		for _, b := range g.Bags {
			for _, c := range b {
				curMin := min[c.Color]

				if c.Amount > curMin {
					min[c.Color] = c.Amount
				}
			}
		}

		acc += min[RED] * min[GREEN] * min[BLUE]
	}

	fmt.Printf("Part Two: %d\n", acc)

	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "2",
			RunE: day,
		},
		"2023",
		"2",
	)
}
