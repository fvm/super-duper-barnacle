package day02

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/sirupsen/logrus"
)

// Input ...
type Input struct {
	Filename     string
	TargetTuples []int
}

type Sequences []Sequence

// Sequence
type Sequence struct {
	characterCount characterCount
	tupleCount     tupleCount
	chars          []string
}

// Solve
func (i Input) Solve() error {
	input, err := openInput(i.Filename)
	if err != nil {
		return err
	}
	sequences := parseInputAsSequences(input)
	n, err := sequences.solvePartOne(i.TargetTuples)
	if err != nil {
		return err
	}
	logrus.Infof("Part one: %d", n)
	s, err := sequences.solvePartTwo()
	if err != nil {
		return err
	}
	logrus.Infof("Part two: %s", s)

	return nil
}

type characterCount struct {
	count   map[string]int
	counted bool
}

type tupleCount struct {
	count   map[int]int
	counted bool
}

func (sequences Sequences) solvePartOne(targets []int) (int, error) {
	hash := 1
	for _, targetTuple := range targets {
		count := 0 // This can go so much faster with channels and goroutines
		for k, sequence := range sequences {
			sequence.countCharacterOccurrences()
			if sequence.hasTuple(targetTuple) {
				count++
			}
			sequences[k] = sequence
		}
		if count != 0 {
			hash *= count
		}
	}
	return hash, nil
}

func (sequences Sequences) solvePartTwo() (string, error) {
	for _, sequence := range sequences {
		needle, err := sequence.findMatchForSequence(sequences)
		if err != nil {
			if err.Error() == "not found" {
				continue
			}
			return "", err
		}
		var result []string
		for key, char := range needle.chars {
			if needle.chars[key] == sequence.chars[key] {
				result = append(result, char)
			}
		}
		return strings.Join(result, ""), nil
	}
	return "", nil
}

func openInput(filename string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func parseInputAsSequences(reader io.Reader) Sequences {
	scanner := bufio.NewScanner(reader)
	var sequences Sequences
	for scanner.Scan() {
		line := scanner.Text()
		sequences = append(sequences, NewSequenceFromString(line))
	}
	return sequences
}

func (s Sequence) findMatchForSequence(haystack []Sequence) (Sequence, error) {
	for _, h := range haystack {
		d, err := hamming(s.chars, h.chars)
		if err != nil {
			logrus.Fatal(err)
		}
		if d == 1 {
			return h, nil
		}
	}
	return Sequence{}, errors.New("not found")
}

func hamming(s1 []string, s2 []string) (int, error) {
	if len(s1) != len(s2) {
		return 0, errors.New("s1 s2 are not the same length")
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		// This is ascii
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff, nil
}

func NewSequenceFromString(str string) Sequence {
	return Sequence{
		characterCount: characterCount{
			count:   map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0},
			counted: false,
		},
		tupleCount: tupleCount{
			count:   map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0, 9: 0, 10: 0, 11: 0, 12: 0, 13: 0, 14: 0, 15: 0, 16: 0, 17: 0, 18: 0, 19: 0, 20: 0, 21: 0, 22: 0, 23: 0, 24: 0, 25: 0, 26: 0},
			counted: false,
		},
		chars: strings.Split(str, ""),
	}
}

func (s *Sequence) hasTuple(n int) bool {
	if s.characterCount.counted == false {
		s.countCharacterOccurrences()
	}
	for _, value := range s.characterCount.count {
		if value == n {
			return true
		}
	}
	return false
}

func (s *Sequence) howManyTuplesOf(n int) int {
	s.countCharacterOccurrences()
	s.countTuples()
	return s.tupleCount.count[n]
}

func (s *Sequence) countTuples() {
	if s.tupleCount.counted == true {
		return
	}
	for _, value := range s.characterCount.count {
		s.tupleCount.count[value]++
	}
	s.tupleCount.counted = true
	return
}

func (s *Sequence) countCharacterOccurrences() {
	if s.characterCount.counted == true {
		return
	}
	for _, value := range s.chars {
		s.characterCount.count[value]++
	}
	s.characterCount.counted = true
	return
}
