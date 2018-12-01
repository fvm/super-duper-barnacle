package day01

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type Input struct {
	Filename string
}

func (d Input) Solve() error {
	if n, err := solve(d); err != nil {
		return err
	} else {
		logrus.Infof("Frequency: %d", n)
	}
	return nil
}

func solve(d Input) (sum int, err error) {
	numbers, err := readLinesFromInputFile(d.Filename)
	if err != nil {
		return sum, err
	}
	seen := make(map[int]int)
	for i := 0; ; i++ {
		for _, v := range numbers {
			sum += v
			seen[sum] += 1
			if v, ok := seen[sum]; ok {
				if v == 2 {
					logrus.Infof("Seen frequency: %d for the second time", sum)
					return sum, nil
				}
			} else {
				seen[sum] = 0
			}
		}
	}
}

func readLinesFromInputFile(filename string) ([]int, error) {
	f, err := os.Open(filename)
	defer func() {
		if err := f.Close(); err != nil {
			logrus.Fatal(err)
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
