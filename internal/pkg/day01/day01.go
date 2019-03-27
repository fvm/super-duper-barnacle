package day01

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Input serves as a type for main to pass stuff to
type Input struct {
	Filename string
}

// Solve for the given Input
func (d Input) Solve() error {
	n, err := solve(d)
	if err != nil {
		return err
	}
	logrus.Infof("Seen frequency: %d for the second time", n)
	return nil
}

func solve(d Input) (int, error) {
	r, err := openInput(d.Filename)
	if err != nil {
		return 0, err
	}
	numbers, err := scanNumbers(r)
	if err != nil {
		return 0, err
	}
	n, err := solveForSeries(numbers)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func solveForSeries(s []int) (int, error) {
	sum := 0
	seen := make(map[int]int)
	for i := 0; ; i++ {
		for _, v := range s {
			sum += v
			seen[sum]++
			if v, ok := seen[sum]; ok {
				if v == 2 {
					return sum, nil
				}
			} else {
				seen[sum] = 0
			}
		}
	}
}

func openInput(filename string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func scanNumbers(r io.Reader) ([]int, error) {
	var numbers []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, nil
}
