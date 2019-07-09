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
	"github.com/orn688/advent-of-code-2018/internal/day09"
	"github.com/orn688/advent-of-code-2018/internal/day10"
	"github.com/orn688/advent-of-code-2018/internal/day11"
	"github.com/orn688/advent-of-code-2018/internal/day12"
	"github.com/orn688/advent-of-code-2018/internal/day13"
	"github.com/orn688/advent-of-code-2018/internal/day14"
)

type dayfunc func(string) (string, error)

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

	fun, err := funcForDay(day, part2)
	if err != nil {
		return err
	}
	result, err := fun(input)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func funcForDay(day int, part2 bool) (fun dayfunc, err error) {
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
	case 9:
		fun = day09.Part1
		if part2 {
			fun = day09.Part2
		}
	case 10:
		fun = day10.Part1
		if part2 {
			fun = day10.Part2
		}
	case 11:
		fun = day11.Part1
		if part2 {
			fun = day11.Part2
		}
	case 12:
		fun = day12.Part1
		if part2 {
			fun = day12.Part2
		}
	case 13:
		fun = day13.Part1
		if part2 {
			fun = day13.Part2
		}
	case 14:
		fun = day14.Part1
		if part2 {
			fun = day14.Part2
		}
	default:
		err = fmt.Errorf("day %d is not implemented", day)
	}
	return fun, err
}
