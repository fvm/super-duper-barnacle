package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	n, err := solvePuzzle();
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("Result: %d", n)
}

func solvePuzzle() (n int, err error) {
	numbers, err := readLinesFromInputFile()
	var sum int
	for k, v := range numbers {
		sum += v
		logrus.Infof("n: %d, v:%d, s: %d", k, v, sum)
	}
	return sum, nil
}

func readLinesFromInputFile() ([]int, error) {
	f, err := os.Open("input.txt")
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return nil, err
	}
	var numbers []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if number, err := strconv.Atoi(line); err != nil {
			continue
		} else {
			numbers = append(numbers, number)
		}
	}
	return numbers, nil
}
