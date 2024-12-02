package main

import (
	"fmt"

	"aoc2024/solutions/day01"
	"aoc2024/solutions/day02"

	"github.com/alecthomas/kong"
)

type CLI struct {
	Day    int    `help:"The day of the challenge to execute" required:""`
	Part   int    `help:"The part of the challenge to execute" required:""`
	Inputs string `help:"Path to the inputs file for the challenge" required:""`
}

func main() {
	var cli CLI
	ctx := kong.Parse(&cli,
		kong.Name("aoc-2024-runner"),
		kong.Description("Advent of Code 2024 runner"),
		kong.UsageOnError(),
	)

	// Run
	err := run(cli.Day, cli.Part, cli.Inputs)
	if err != nil {
		fmt.Printf("Error: %v", err)
		ctx.Exit(1)
	}

	ctx.Exit(0)
}

func run(day, part int, inputs string) error {
	switch day {
	case 1:
		switch part {
		case 1:
			day01.Part1(inputs)
		case 2:
			day01.Part2(inputs)
		default:
			return fmt.Errorf("each day only have parts 1 and 2")
		}
	case 2:
		switch part {
		case 1:
			day02.Part1(inputs)
		case 2:
			day02.Part2(inputs)
		default:
			return fmt.Errorf("each day only have parts 1 and 2")
		}
	default:
		return fmt.Errorf("day %d is not implemented", day)
	}
	return nil
}
