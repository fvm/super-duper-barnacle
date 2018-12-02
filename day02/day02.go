package day02

import (
	"bufio"
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"strings"
)

type Input struct {
	Filename     string
	TargetTuples []int
}

type Sequences []Sequence

type characterCount struct {
	count   map[string]int
	counted bool
}

type tupleCount struct {
	count   map[int]int
	counted bool
}

type Sequence struct {
	characterCount characterCount
	tupleCount     tupleCount
	chars          []string
}

func (i Input) Solve() error {
	n, err := i.solvePartOne()
	if err != nil {
		return err
	}
	logrus.Infof("Found %d", n)
	return nil
}
func (i Input) solvePartOne() (int, error) {
	input, err := openInput(i.Filename)
	if err != nil {
		return 0, err
	}
	hash := 1
	sequences := parseInput(input)
	for _, targetTuple := range i.TargetTuples {
		count := 0 // This can go so much faster with channels and goroutines
		for _, sequence := range sequences {
			if sequence.hasTuple(targetTuple) {
				count++
			}
		}
		if count != 0 {
			hash *= count
		}
	}
	return hash, nil
}
func openInput(filename string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func parseInput(reader io.Reader) Sequences {
	scanner := bufio.NewScanner(reader)
	var sequences Sequences
	for scanner.Scan() {
		line := scanner.Text()
		sequences = append(sequences, NewSequenceFromString(line))
	}
	return sequences
}

func NewSequenceFromString(str string) Sequence {
	return Sequence{
		characterCount: characterCount{
			count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0,},
			counted: false,
		},
		tupleCount: tupleCount{
			count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
			counted: false,
		},
		chars: strings.Split(str, ""),
	}
}

func NewSequence() Sequence {
	return Sequence{
		characterCount: characterCount{
			count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0,},
			counted: false,
		},
		tupleCount: tupleCount{
			count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
			counted: false,
		},
		chars: []string{},
	}
}

func (s Sequence) hasTuple(n int) bool {
	s.countCharacterOccurrences()
	for _, value := range s.characterCount.count {
		if value == n {
			return true
		}
	}
	return false
}

func (s Sequence) howManyTuplesOf(n int) int {
	return s.countCharacterOccurrences().countTuples().tupleCount.count[n]
}

func (s Sequence) countTuples() Sequence {
	if s.tupleCount.counted == true {
		return s
	}
	for _, value := range s.characterCount.count {
		s.tupleCount.count[value]++
	}
	s.tupleCount.counted = true
	return s
}

func (s Sequence) countCharacterOccurrences() Sequence {
	if s.characterCount.counted == true {
		return s
	}
	for _, value := range s.chars {
		s.characterCount.count[value]++
	}
	s.characterCount.counted = true
	return s
}
