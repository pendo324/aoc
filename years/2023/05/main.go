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
	in, err := os.ReadFile(filepath.Join("years", "2023", "05", "input"))
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}

	split := strings.Split(string(in), "\nseed-to-soil map:\n")
	s, _ := strings.CutPrefix(split[0], "seeds: ")
	seeds := strings.TrimSpace(s)
	maps := parseMaps(split[1])

	if err := part1(seeds, maps); err != nil {
		return fmt.Errorf("failed to run part1: %w", err)
	}

	if err := part2(seeds, maps); err != nil {
		return fmt.Errorf("failed to run part2: %w", err)
	}

	return nil
}

func part1(seeds string, maps [][][3]int) error {
	seedsWithLen := ""
	for i, s := range strings.Split(seeds, " ") {
		if i != 0 {
			seedsWithLen += " "
		}
		seedsWithLen += s
		seedsWithLen += " 1"
	}

	fmt.Printf("Part One: %v\n", findMinSeed(seedsWithLen, maps))

	return nil
}

func part2(seeds string, maps [][][3]int) error {
	fmt.Printf("Part Two: %v\n", findMinSeed(seeds, maps))

	return nil
}

func parseMaps(in string) [][][3]int {
	s1 := strings.Split(in, "\nsoil-to-fertilizer map:\n")
	soilToFert := s1[0]
	s2 := strings.Split(s1[1], "\nfertilizer-to-water map:\n")
	seedsToSoil := s2[0]
	s3 := strings.Split(s2[1], "\nwater-to-light map:\n")
	fertToWater := s3[0]
	s4 := strings.Split(s3[1], "\nlight-to-temperature map:\n")
	waterToLight := s4[0]
	s5 := strings.Split(s4[1], "\ntemperature-to-humidity map:\n")
	lightToTemp := s5[0]
	s6 := strings.Split(s5[1], "\nhumidity-to-location map:\n")
	tempToHumidity := s6[0]
	humidityToLocation := s6[1]

	mapStrings := [7]string{
		soilToFert,
		seedsToSoil,
		fertToWater,
		waterToLight,
		lightToTemp,
		tempToHumidity,
		humidityToLocation,
	}

	// maps are not a good idea for things that take billions of lookups, use a slice.
	maps := [][][3]int{}
	for _, sMap := range mapStrings {
		ls := strings.Split(strings.TrimSpace(sMap), "\n")

		out := [][3]int{}
		for _, l := range ls {
			s := strings.Split(l, " ")
			dest, _ := strconv.Atoi(s[0])
			source, _ := strconv.Atoi(s[1])
			length, _ := strconv.Atoi(s[2])

			out = append(out, [3]int{source, dest, length})
		}

		maps = append(maps, out)
	}

	return maps
}

func findMinSeed(seeds string, maps [][][3]int) int {
	ss := strings.Split(seeds, " ")
	minLoc := 0
	for i := 0; i < len(ss)-1; i = i + 2 {
		seed, _ := strconv.Atoi(ss[i])
		r, _ := strconv.Atoi(ss[i+1])
		for j := seed; j < seed+r; j++ {
			currentLoc := j
			for _, m := range maps {
				for _, l := range m {
					source := l[0]
					dest := l[1]
					length := l[2]
					if currentLoc < source+length && currentLoc >= source {
						currentLoc = (dest + length) - (source + length - currentLoc)
						break
					}
				}
			}
			if minLoc == 0 || currentLoc < minLoc {
				minLoc = currentLoc
			}
		}
	}

	return minLoc
}

// Original map solution which did NOT scale.
// func part1(seeds, in string) error {
// 	s1 := strings.Split(in, "\nsoil-to-fertilizer map:\n")
// 	soilToFert := s1[0]
// 	s2 := strings.Split(s1[1], "\nfertilizer-to-water map:\n")
// 	seedsToSoil := s2[0]
// 	s3 := strings.Split(s2[1], "\nwater-to-light map:\n")
// 	fertToWater := s3[0]
// 	s4 := strings.Split(s3[1], "\nlight-to-temperature map:\n")
// 	waterToLight := s4[0]
// 	s5 := strings.Split(s4[1], "\ntemperature-to-humidity map:\n")
// 	lightToTemp := s5[0]
// 	s6 := strings.Split(s5[1], "\nhumidity-to-location map:\n")
// 	tempToHumidity := s6[0]
// 	humidityToLocation := s6[1]

// 	mapStrings := [7]string{
// 		soilToFert,
// 		seedsToSoil,
// 		fertToWater,
// 		waterToLight,
// 		lightToTemp,
// 		tempToHumidity,
// 		humidityToLocation,
// 	}

// 	maps := []map[image.Point]int{}
// 	for _, sMap := range mapStrings {
// 		ls := strings.Split(strings.TrimSpace(sMap), "\n")

// 		out := map[image.Point]int{}
// 		for _, l := range ls {
// 			s := strings.Split(l, " ")
// 			dest, _ := strconv.Atoi(s[0])
// 			source, _ := strconv.Atoi(s[1])
// 			length, _ := strconv.Atoi(s[2])

// 			out[image.Pt(source, dest)] = length
// 		}

// 		maps = append(maps, out)
// 	}

// 	minLoc := 0
// 	for _, s := range strings.Split(seeds, " ") {
// 		seed, _ := strconv.Atoi(strings.TrimSpace(s))
// 		lookups := [8]int{seed}
// 		for i, m := range maps {
// 			for p, l := range m {
// 				if lookups[i] < p.X+l && lookups[i] >= p.X {
// 					lookups[i+1] = (p.Y + l) - (p.X + l - lookups[i])
// 					break
// 				}
// 			}
// 			if lookups[i+1] == 0 {
// 				lookups[i+1] = lookups[i]
// 			}
// 		}
// 		if minLoc == 0 || lookups[7] < minLoc {
// 			minLoc = lookups[7]
// 		}
// 	}

// 	fmt.Printf("Part One: %v\n", minLoc)

// 	return nil
// }

func init() {
	cmd.RegisterYearDay(
		cmd.RunDayCmd,
		&cobra.Command{
			Use:  "5",
			RunE: day,
		},
		"2023",
		"5",
	)
}
