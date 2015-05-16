package hex

import (
	"math"
)

type Orientation struct {
	f0, f1, f2, f3 float64
	b0, b1, b2, b3 float64
	startAngle     float64
}

var layoutPoint Orientation = Orientation{
	math.Sqrt(3.0), math.Sqrt(3.0) / 2.0, 0.0, 3.0 / 2.0,
	math.Sqrt(3.0) / 3.0, -1.0 / 3.0, 0.0, 2.0 / 3.0,
	0.5,
}

var layoutFlat Orientation = Orientation{
	3.0 / 2.0, 0.0, math.Sqrt(3.0) / 2.0, math.Sqrt(3.0),
	2.0 / 3.0, 0.0, -1.0 / 3.0, math.Sqrt(3.0) / 3.0,
	0.0,
}

type Point struct {
	x, y float64
}

type Layout struct {
	orientation Orientation
	size        Point
	offset      Point
}

func (layout Layout) ToPixel(h Hex) Point {
	M := layout.orientation
	x := (M.f0*float64(h.q) + M.f1*float64(h.r)) * float64(layout.size.x)
	y := (M.f2*float64(h.q) + M.f3*float64(h.r)) * float64(layout.size.y)

	return Point{
		x + layout.offset.x,
		y + layout.offset.y,
	}
}

func (layout Layout) ToHex(p Point) FractionalHex {
	M := layout.orientation
	pt := Point{
		x: (p.x - layout.offset.x) / layout.size.x,
		y: (p.y - layout.offset.y) / layout.size.y,
	}
	q := M.b0*pt.x + M.b1*pt.y
	r := M.b2*pt.x + M.b3*pt.y
	return FractionalHex{
		q, r, -q - r,
	}
}

func (layout Layout) HexCornerOffset(corner int) Point {
	size := layout.size
	angle := 2.0 * math.Pi * (float64(corner) + layout.orientation.startAngle) / 6
	return Point{size.x * math.Cos(angle), size.y * math.Sin(angle)}
}

func (layout Layout) PolygonCorners(hex Hex) []Point {
	corners := make([]Point, 6)
	center := layout.ToPixel(hex)

	for i := 0; i < 6; i++ {
		offset := layout.HexCornerOffset(i)
		corners[i] = Point{center.x + offset.x, center.y + offset.y}
	}

	return corners
}
