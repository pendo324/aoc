package day

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pendo324/aoc/cmd"
	"github.com/spf13/cobra"
)

func day(cmd *cobra.Command, _ []string) error {
	in, err := os.ReadFile(filepath.Join("years", "2023", "06", "input"))
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
	re := regexp.MustCompile(`\w*:\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	m := re.FindAllStringSubmatch(in, -1)

	times := m[0][1:]
	distances := m[1][1:]

	wins := [4]int{}
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		for j := 0; j < time; j++ {
			if (time-j)*j > distance {
				wins[i]++
			}
		}
	}

	acc := 1
	for _, win := range wins {
		acc *= win
	}

	fmt.Printf("Part One: %d\n", acc)

	return nil
}

func part2(in string) error {
	re := regexp.MustCompile(`\w*:\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	m := re.FindAllStringSubmatch(in, -1)

	time := strings.Join(m[0][1:], "")
	distance := strings.Join(m[1][1:], "")

	t, _ := strconv.Atoi(time)
	d, _ := strconv.Atoi(distance)

	wins := 0
	for i := 0; i < t; i++ {
		if (t-i)*i > d {
			wins++
		}
	}

	fmt.Printf("Part Two: %d\n", wins)

	return nil
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "6",
			RunE: day,
		},
		"2023",
		"6",
	)
}
