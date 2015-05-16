package hex

import (
	"errors"
	"math"
)

type Hex struct {
	q, r, s int
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

func (self Hex) Multiply(k int) Hex {
	return Hex{
		q: self.q * k,
		r: self.r * k,
		s: self.s * k,
	}
}

func (self Hex) Length() int {
	return int((math.Abs(float64(self.q)) + math.Abs(float64(self.r)) + math.Abs(float64(self.s))) / 2)
}

func (self Hex) Distance(other Hex) int {
	return self.Substract(other).Length()
}

func (self Hex) Lerp(other Hex, t float64) FractionalHex {
	return FractionalHex{
		float64(self.q+(other.q-self.q)) * t,
		float64(self.r+(other.r-self.r)) * t,
		float64(self.s+(other.s-self.s)) * t,
	}
}

func (self Hex) LineDraw(other Hex) []Hex {
	N := self.Distance(other)
	result := make([]Hex, N+1)
	step := 1.0 / math.Max(float64(N), float64(1))

	for i := 0; i <= N; i++ {
		result[i] = self.Lerp(other, step*float64(i)).Round()
	}
	return result
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
	q := int(self.q)
	r := int(self.r)
	s := int(self.s)

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
