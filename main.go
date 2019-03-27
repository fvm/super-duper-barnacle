package main

import (
	"github.com/fvm/super-duper-barnacle-AoC-2018/internal/pkg/day01"
	"github.com/fvm/super-duper-barnacle-AoC-2018/internal/pkg/day02"
	"github.com/sirupsen/logrus"
)

//noinspection GoUnusedExportedType (it's used)
type Solver interface {
	Solve() error
}

func main() {
	day01 := day01.Input{
		Filename: "./day01/input.txt",
	}
	err := day01.Solve()
	if err != nil {
		logrus.Error(err)
	}
	day02 := day02.Input{
		Filename:     "./day02/input.txt",
		TargetTuples: []int{2, 3},
	}
	err = day02.Solve()
	if err != nil {
		logrus.Error(err)
	}
}
