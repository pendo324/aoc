package day

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "04", "input"))
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
	acc := 0
	for _, l := range strings.Split(in, "\n") {
		m := regexp.MustCompile(
			`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+\|` +
				`\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+` +
				`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+` +
				`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`,
		).FindStringSubmatch(l)
		w := m[1:11]
		h := m[11:]

		matches := 0
		for _, i := range h {
			for _, j := range w {
				if i == j {
					matches++
				}
			}
		}

		acc += int(math.Pow(2, float64(matches-1)))
	}

	fmt.Printf("Part One: %d\n", acc)

	return nil
}

func part2(in string) error {
	rows := strings.Split(in, "\n")
	r := make([]int, len(rows))
	for i := range r {
		r[i] = 1
	}

	for id, v := range r {
		row := rows[id]
		m := regexp.MustCompile(
			`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+\|` +
				`\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+` +
				`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+` +
				`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`,
		).FindStringSubmatch(row)
		w := m[1:11]
		h := m[11:]

		for i := 0; i < v; i++ {
			matches := 0
			for _, j := range h {
				for _, n := range w {
					if j == n {
						r[id+1+matches] += 1
						matches++
					}
				}
			}
		}
	}

	acc := 0
	for _, m := range r {
		acc += m
	}

	fmt.Printf("Part Two: %d\n", acc)

	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "4",
			RunE: day,
		},
		"2023",
		"4",
	)
}
