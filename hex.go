package hex

import (
	"errors"
	"math"
)

type Hex struct {
	q, r, s int64
}

var directions = [6]Hex{
	Hex{1, 0, -1}, Hex{1, -1, 0}, Hex{0, -1, 1},
	Hex{-1, 0, 1}, Hex{-1, 1, 0}, Hex{0, 1, -1},
}

func (self Hex) Equal(other Hex) bool {
	return self.q == other.q && self.r == other.r && self.s == other.s
}

func (self Hex) Add(other Hex) Hex {
	return Hex{
		q: self.q + other.q,
		r: self.r + other.r,
		s: self.s + other.s,
	}
}

func (self Hex) Substract(other Hex) Hex {
	return Hex{
		q: self.q - other.q,
		r: self.r - other.r,
		s: self.s - other.s,
	}
}

func (self Hex) Multiply(k int64) Hex {
	return Hex{
		q: self.q * k,
		r: self.r * k,
		s: self.s * k,
	}
}

func (self Hex) Length() int64 {
	return int64((math.Abs(float64(self.q)) + math.Abs(float64(self.r)) + math.Abs(float64(self.s))) / 2)
}

func (self Hex) Distance(other Hex) int64 {
	return self.Substract(other).Length()
}

func Direction(direction int) (Hex, error) {
	if 0 <= direction && direction < 6 {
		return Hex{}, errors.New("out of directions")
	}
	return directions[direction], nil
}

type FractionalHex struct {
	q, r, s float64
}

func (self FractionalHex) Round() Hex {
	q := int64(self.q)
	r := int64(self.r)
	s := int64(self.s)

	qDiff := math.Abs(float64(q) - self.q)
	rDiff := math.Abs(float64(r) - self.r)
	sDiff := math.Abs(float64(s) - self.s)

	if qDiff > rDiff && qDiff > sDiff {
		q = -r - s
	} else if rDiff > sDiff {
		r = -q - s
	} else {
		s = -q - r
	}

	return Hex{q, r, s}
}
