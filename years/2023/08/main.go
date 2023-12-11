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
	in, err := os.ReadFile(filepath.Join("years", "2023", "08", "input"))
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
	instructions := ""
	ns := make([][2]int, 909091)

	cur := [2]int{}
	for i, l := range strings.Split(in, "\n") {
		if i == 0 {
			instructions = l
		} else if l != "" {
			re := regexp.MustCompile(`(\w{3})\s=\s\((\w{3}),\s(\w{3})\)`)
			m := re.FindAllStringSubmatch(l, -1)
			for _, match := range m {
				keys := [3]int{}
				for keyIdx, key := range match {
					if keyIdx == 0 {
						continue
					}
					sString := ""
					for _, r := range key {
						sString += fmt.Sprintf("%d", r)
					}
					source, _ := strconv.Atoi(sString)
					keys[keyIdx-1] = source
				}
				ns[keys[0]] = [2]int{keys[1], keys[2]}
			}
		}
	}

	cur = ns[656565]
	iterations := 1
	for i := 0; i < len(instructions); {
		r := instructions[i]
		if r == 'L' {
			if cur[0] == 909090 {
				break
			}
			cur = ns[cur[0]]
		} else if r == 'R' {
			if cur[1] == 909090 {
				break
			}
			cur = ns[cur[1]]
		}

		if i == len(instructions)-1 {
			i = -1
		}

		i++
		iterations++
	}

	fmt.Printf("Part One: %d\n", iterations)

	return nil
}

func part2(in string) error {
	instructions := ""
	ns := make([][2]int, 909091)

	startNodes := [][2]int{}
	for i, l := range strings.Split(in, "\n") {
		if i == 0 {
			instructions = l
		} else if l != "" {
			re := regexp.MustCompile(`(\w{3})\s=\s\((\w{3}),\s(\w{3})\)`)
			m := re.FindAllStringSubmatch(l, -1)
			for _, match := range m {
				keys := [3]int{}
				for keyIdx, key := range match {
					if keyIdx == 0 {
						continue
					}
					sString := ""
					for _, r := range key {
						sString += fmt.Sprintf("%d", r)
					}
					source, _ := strconv.Atoi(sString)
					keys[keyIdx-1] = source
				}
				if strings.HasSuffix(match[1], "A") {
					startNodes = append(startNodes, [2]int{keys[1], keys[2]})
				}
				ns[keys[0]] = [2]int{keys[1], keys[2]}
			}
		}
	}

	cur := make([][2]int, len(startNodes))
	zs := []int{}

	for j := 0; j < len(cur); j++ {
		iterations := 1
		for i := 0; i < len(instructions); {
			r := instructions[i]
			if r == 'L' {
				if strings.HasSuffix(fmt.Sprintf("%d", startNodes[j][0]), "90") {
					zs = append(zs, iterations)
					break
				}
				startNodes[j] = ns[startNodes[j][0]]
			} else if r == 'R' {
				if strings.HasSuffix(fmt.Sprintf("%d", startNodes[j][1]), "90") {
					zs = append(zs, iterations)
					break
				}
				startNodes[j] = ns[startNodes[j][1]]
			}

			if i == len(instructions)-1 {
				i = -1
			}
			i++
			iterations++
		}
	}

	fmt.Printf("Part Two: %d\n", LCM(zs[0], zs[1], zs[1:]...))

	return nil
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "8",
			RunE: day,
		},
		"2023",
		"8",
	)
}
