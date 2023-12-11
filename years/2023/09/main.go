package day

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "09", "input"))
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

func reduce(count int, diffs [][]int) int {
	out := append(diffs, make([]int, len(diffs[count-1])-1))
	zCount := 0
	for i := 0; i < len(out[count-1])-1; i++ {
		diff := out[count-1][i+1] - out[count-1][i]
		out[count][i] = diff
		if diff == 0 {
			zCount++
		}
	}

	if zCount != len(out[count]) {
		return reduce(count+1, out)
	} else {
		newVal := 0
		for i := len(out) - 1; i > 0; i-- {
			lastCur := out[i][len(out[i])-1]
			lastPrev := out[i-1][len(out[i-1])-1]
			newVal = lastCur + lastPrev
			out[i-1] = append(out[i-1], lastCur+lastPrev)
		}

		return newVal
	}
}

func part1(in string) error {
	acc := 0
	for _, l := range strings.Split(in, "\n") {
		numString := strings.Split(l, " ")
		nums := []int{}
		for _, s := range numString {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		acc += reduce(1, [][]int{nums})
	}
	fmt.Printf("Part One: %d\n", acc)
	return nil
}

func part2(in string) error {
	acc := 0
	for _, l := range strings.Split(in, "\n") {
		numString := strings.Split(l, " ")
		nums := []int{}
		for _, s := range numString {
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}
		slices.Reverse(nums)
		acc += reduce(1, [][]int{nums})
	}
	fmt.Printf("Part Two: %d\n", acc)
	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "9",
			RunE: day,
		},
		"2023",
		"9",
	)
}
