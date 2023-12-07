package main

import (
	"fmt"
	"os"

	"github.com/pendo324/aoc/cmd"
	_ "github.com/pendo324/aoc/years"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Running aoc")
	cobra.OnInitialize(initConfig)

	cmd.RootCmd.AddCommand(cmd.RunDayCmd)

	cmd.NewCreateDayCmd(cmd.RootCmd)

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./../")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		os.Exit(1)
	}
}
