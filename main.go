package main

import (
	"github.com/sirupsen/logrus"

	"github.com/fvm/super-duper-barnacle/internal/day01"
	"github.com/fvm/super-duper-barnacle/internal/day02"
	"github.com/fvm/super-duper-barnacle/internal/day03"
)

func main() {
	day01input := day01.Input{
		Filename: "./day01input/input.txt",
	}
	err := day01input.Solve()
	if err != nil {
		logrus.Error(err)
	}

	day02input := day02.Input{
		Filename:     "./day02input/input.txt",
		TargetTuples: []int{2, 3},
	}
	err = day02input.Solve()
	if err != nil {
		logrus.Error(err)
	}

	day03input := day03.Input{}
	err = day03input.Solve()
	if err != nil {
		logrus.Error(err)
	}
}
