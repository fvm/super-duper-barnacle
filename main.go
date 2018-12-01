package main

import (
	"github.com/fvm/super-duper-barnacle-AoC-2018/day01"
	"github.com/sirupsen/logrus"
)

//noinspection GoUnusedExportedType (it's used)
type Solver interface {
	Solve() error
}

func main() {
	in := day01.Input{
		Filename: "./day01/input.txt",
	}
	err := in.Solve()
	if err != nil {
		logrus.Error(err)
	}
}
