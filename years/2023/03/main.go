package day

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "03", "input"))
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}

	rows := strings.Split(string(in), "\n")
	part1(rows)
	part2(rows)

	return nil
}

func setup(rows []string, f func(rune) bool) map[image.Point][]int {
	partNumbers := map[image.Point]rune{}
	adj := map[image.Point][]int{}
	rowLen := 0

	for i := 0; i < len(rows); i++ {
		row := rows[i]
		for j, col := range row {
			rowLen = len(row)
			partNumbers[image.Pt(i, j)] = col
		}
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < rowLen; j++ {
			p := image.Pt(i, j)
			n := partNumbers[p]
			if unicode.IsDigit(n) {
				numAcc := string(n)
				numPoints := []image.Point{image.Pt(i, j)}
				ps := map[image.Point]bool{}

				for k := j + 1; ; k++ {
					pp := p.Add(image.Pt(0, k-j))
					nn := partNumbers[pp]
					if unicode.IsDigit(nn) {
						numAcc += string(nn)
						numPoints = append(numPoints, image.Pt(i, k))
					} else {
						j = k
						break
					}
				}

				num, _ := strconv.Atoi(numAcc)

				adjacentPoints := []image.Point{
					{-1, -1},
					{-1, 0},
					{-1, 1},
					{0, -1},
					{0, 1},
					{1, -1},
					{1, 0},
					{1, 1},
				}

				for _, k := range numPoints {
					for _, p := range adjacentPoints {
						newPoint := k.Add(p)
						if f(partNumbers[newPoint]) {
							ps[newPoint] = true
						}
					}
				}

				for p := range ps {
					adj[p] = append(adj[p], num)
				}
			}
		}
	}

	return adj
}

func part1(rows []string) error {
	adj := setup(rows, func(r rune) bool {
		return r != '.' && !unicode.IsDigit(r) && r != 0
	})

	acc := 0
	for _, nums := range adj {
		for _, n := range nums {
			acc += n
		}
	}

	fmt.Printf("Part One: %d\n", acc)

	return nil
}

func part2(rows []string) error {
	adj := setup(rows, func(r rune) bool {
		return r == '*'
	})

	acc := 0
	for _, nums := range adj {
		if len(nums) == 2 {
			acc += nums[0] * nums[1]
		}
	}

	fmt.Printf("Part Two: %d\n", acc)

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
