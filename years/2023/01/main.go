package day

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile("years\\2023\\01\\input")
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

func part1(in string) error {
	res := []int{}
	for _, l := range strings.Split(in, "\n") {
		first := ""
		second := ""

		for _, c := range l {
			if unicode.IsDigit(c) {
				if first == "" {
					first = string(c)
				} else {
					second = string(c)
				}
			}
		}
		acc := first
		if second == "" {
			acc += first
		} else {
			acc += second
		}

		o, err := strconv.Atoi(acc)
		if err != nil {
			return fmt.Errorf("failed to convert to int: %w", err)
		}

		res = append(res, o)
	}

	out := 0
	for _, i := range res {
		out += i
	}

	fmt.Printf("Part One: %d\n", out)

	return nil
}

func part2(in string) error {
	res := []int{}
	r := strings.NewReplacer(
		"one", "o1e",
		"two", "t2o",
		"three", "t3e",
		"four", "fo4r",
		"five", "fi5e",
		"six", "s6x",
		"seven", "se7en",
		"eight", "ei8ht",
		"nine", "ni9e",
	)
	for _, l := range strings.Split(in, "\n") {
		first := ""
		second := ""

		rp := r.Replace(r.Replace(l))

		for i := 0; i < len(rp); i++ {
			c := rune(rp[i])
			if unicode.IsDigit(c) {
				if first == "" {
					first = string(c)
				} else {
					second = string(c)
				}
			}
		}
		acc := first
		if second == "" {
			acc += first
		} else {
			acc += second
		}

		o, err := strconv.Atoi(acc)
		if err != nil {
			return fmt.Errorf("failed to convert to int: %w", err)
		}

		res = append(res, o)
	}

	out := 0
	for _, i := range res {
		out += i
	}

	fmt.Printf("Part Two: %d\n", out)

	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "1",
			RunE: day,
		},
		"2023",
		"1",
	)
}
