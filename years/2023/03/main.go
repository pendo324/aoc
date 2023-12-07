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

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "03", "input"))
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

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func intPointer(b byte) *int {
	i := int(b)
	return &i
}

func part1(in string) error {
	rows := strings.Split(in, "\n")
	partNumbers := make([][]*int, len(rows))
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if i == 0 {
			partNumbers[i] = make([]*int, len(row))
		}
		if i < len(rows)-1 {
			partNumbers[i+1] = make([]*int, len(row))
		}
		for j, col := range row {
			if col != '.' && !isDigit(col) {
				// left
				if j > 0 && isDigit(rune(row[j-1])) {
					partNumbers[i][j-1] = intPointer(row[j-1] - '0')
				}
				// right
				if j < len(row)-1 && isDigit(rune(row[j+1])) {
					partNumbers[i][j+1] = intPointer(row[j+1] - '0')
				}
				// above
				if isDigit(rune(rows[i-1][j])) {
					partNumbers[i-1][j] = intPointer(rows[i-1][j] - '0')
				}
				// top left
				if j > 0 && isDigit(rune(rows[i-1][j-1])) {
					partNumbers[i-1][j-1] = intPointer(rows[i-1][j-1] - '0')
				}
				// top right
				if j < len(row)-1 && isDigit(rune(rows[i-1][j+1])) {
					partNumbers[i-1][j+1] = intPointer(rows[i-1][j+1] - '0')
				}
				if i < len(row)-1 {
					// below
					if isDigit(rune(rows[i+1][j])) {
						partNumbers[i+1][j] = intPointer(rows[i+1][j] - '0')
					}
					// bottom left
					if j > 0 && isDigit(rune(rows[i+1][j-1])) {
						partNumbers[i+1][j-1] = intPointer(rows[i+1][j-1] - '0')
					}
					// bottom right
					if j < len(row)-1 && isDigit(rune(rows[i+1][j+1])) {
						partNumbers[i+1][j+1] = intPointer(rows[i+1][j+1] - '0')
					}
				}
			}
		}
	}

	for i, row := range partNumbers {
		for j := range row {
			if row[j] != nil {
				if j < len(row)-1 {
					expand(rows[i], partNumbers, i, j+1)
				}
				if j > 0 {
					expand(rows[i], partNumbers, i, j-1)
				}
			}
		}
	}

	acc := 0
	for _, row := range partNumbers {
		s := ""
		for _, col := range row {
			if col == nil {
				n, _ := strconv.Atoi(s)
				acc += n
				s = ""
			} else {
				s += strconv.Itoa(*col)
			}
		}
		n, _ := strconv.Atoi(s)
		acc += n
	}

	fmt.Printf("Part One: %d\n", acc)

	return nil
}

func expand(row string, arr [][]*int, i, j int) {
	if j > 0 {
		if arr[i][j] == nil {
			if isDigit(rune(row[j])) {
				arr[i][j] = intPointer(row[j] - '0')
				expand(row, arr, i, j-1)
			}
		}
	}
	if j < len(arr[i])-1 {
		if arr[i][j] == nil {
			if isDigit(rune(row[j])) {
				arr[i][j] = intPointer(row[j] - '0')
				expand(row, arr, i, j+1)
			}
		}
	}
}

func part2(in string) error {
	fmt.Printf("Part Two: %d\n", 0)

	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "3",
			RunE: day,
		},
		"2023",
		"3",
	)
}
