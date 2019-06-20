package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"

	"github.com/orn688/advent-of-code-2018/internal/client"
	"github.com/orn688/advent-of-code-2018/internal/day01"
	"github.com/orn688/advent-of-code-2018/internal/day02"
	"github.com/orn688/advent-of-code-2018/internal/day03"
	"github.com/orn688/advent-of-code-2018/internal/day04"
	"github.com/orn688/advent-of-code-2018/internal/day05"
	"github.com/orn688/advent-of-code-2018/internal/day06"
	"github.com/orn688/advent-of-code-2018/internal/day07"
	"github.com/orn688/advent-of-code-2018/internal/day08"
)

func main() {
	app := cli.NewApp()
	app.Name = "Advent of Code 2018"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "part2",
		},
	}
	app.Action = action
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func action(context *cli.Context) error {
	if context.NArg() == 0 {
		return errors.New("day must be specified")
	}
	day, err := strconv.Atoi(context.Args().First())
	if err != nil {
		return err
	}
	part2 := context.GlobalBool("part2")
	return runDay(day, part2)
}

func runDay(day int, part2 bool) error {
	input, err := client.GetInput(day)
	if err != nil {
		return err
	}

	var fun func(string) (string, error)

	switch day {
	case 1:
		fun = day01.Part1
		if part2 {
			fun = day01.Part2
		}
	case 2:
		fun = day02.Part1
		if part2 {
			fun = day02.Part2
		}
	case 3:
		fun = day03.Part1
		if part2 {
			fun = day03.Part2
		}
	case 4:
		fun = day04.Part1
		if part2 {
			fun = day04.Part2
		}
	case 5:
		fun = day05.Part1
		if part2 {
			fun = day05.Part2
		}
	case 6:
		fun = day06.Part1
		if part2 {
			fun = day06.Part2
		}
	case 7:
		fun = day07.Part1
		if part2 {
			fun = day07.Part2
		}
	case 8:
		fun = day08.Part1
		if part2 {
			fun = day08.Part2
		}
	default:
		return fmt.Errorf("day %d is not implemented", day)
	}

	result, err := fun(input)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
