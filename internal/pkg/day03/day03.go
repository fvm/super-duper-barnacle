package day03

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
)

type Claim struct {
	// index int
	claim image.Rectangle
}

type Input struct {
}

func (Input) Solve() error {
	return nil
}

type Solver interface {
	Solve()
}

func (c1 Claim) Intersect(c2 Claim) Claim {
	return Claim{c1.claim.Intersect(c2.claim)}
}

func parseInputAsSequences(reader io.Reader) ([]Claim, error) {
	scanner := bufio.NewScanner(reader)
	var claims []Claim
	for scanner.Scan() {
		line := scanner.Text()
		claim, err := newClaim(line)
		if err != nil {
			return nil, err
		}
		claims = append(claims, claim)
	}
	return claims, nil
}

func newClaim(line string) (Claim, error) {
	var idx, x, y, dx, dy int
	// #1 @ 555,891: 18x12
	fString := "#%d @ %d,%d: %dx%d"
	_, err := fmt.Sscanf(line, fString, idx, &x, &y, &dx, &dy)
	if err != nil {
		return Claim{
			image.Rect(0, 0, 0, 0),
		}, err
	}
	return Claim{
		image.Rect(x, y, x+dx, y+dy),
	}, nil
}

func openInput(filename string) (io.Reader, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
