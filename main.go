package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"

	"github.com/orn688/advent-of-code-2018/client"
	"github.com/orn688/advent-of-code-2018/days/day01"
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
		return errors.New("Day must be specified")
	}
	day, err := strconv.Atoi(context.Args().First())
	if err != nil {
		return err
	}
	part2 := context.GlobalBool("part2")
	fmt.Println(part2)
	return runDay(day, part2)
}

func runDay(day int, part2 bool) error {
	input, err := client.GetInput(day)
	if err != nil {
		return err
	}

	fun := func(string) int { return 0 }
	switch day {
	case 1:
		fun = day01.Part1
		if part2 {
			fun = day01.Part2
		}
	default:
		return fmt.Errorf("Day %d is not implemented", day)
	}
	fmt.Println(fun(input))
	return nil
}
