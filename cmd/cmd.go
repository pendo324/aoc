package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "aoc",
	SilenceErrors: true,
	SilenceUsage:  true,
}

var RunDayCmd = &cobra.Command{
	Use: "run",
}

func RegisterYearDay(rootCmd, newDayCmd *cobra.Command, year, day string) {
	var yearCmd *cobra.Command = nil
	for _, c := range rootCmd.Commands() {
		if c.Use == year {
			yearCmd = c
		}
	}

	if yearCmd == nil {
		yearCmd = &cobra.Command{
			Use: year,
		}
		rootCmd.AddCommand(yearCmd)
	}

	var dayCmd *cobra.Command = nil
	for _, c := range yearCmd.Commands() {
		if c.Use == day {
			dayCmd = c
		}
	}

	if dayCmd == nil {
		yearCmd.AddCommand(newDayCmd)
	}
}
