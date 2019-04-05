package day03

import (
	"fmt"
	"image"
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

// Intersect
func (c1 Claim) Intersect(c2 Claim) Claim {
	return Claim{c1.claim.Intersect(c2.claim)}
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
