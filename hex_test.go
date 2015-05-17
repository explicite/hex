package hex

import (
	"testing"
)

func TestEqual(t *testing.T) {
	hex1 := Hex{1.0, 2.0, 3.0}
	hex2 := Hex{4.0, 5.0, 6.0}

	if hex1.Equal(hex2) {
		t.Errorf("hex: %+v and hex: %+v not equal", hex1, hex2)
	}

	hex3 := Hex{1.0, 2.0, 3.0}

	if !hex1.Equal(hex3) {
		t.Errorf("hex: %+v and hex: %+v equal", hex1, hex3)
	}
}

func TestAdd(t *testing.T) {
	hex1 := Hex{11, 33, 92}
	hex2 := Hex{321, 12, 23}
	hex3 := Hex{11 + 321, 33 + 12, 92 + 23}

	if !hex1.Add(hex2).Equal(hex3) {
		t.Errorf("hex: %+v add hex: %+v equal %+v", hex1, hex2, hex3)
	}
}

func TestSubstract(t *testing.T) {
	hex1 := Hex{11, 33, 92}
	hex2 := Hex{321, 12, 23}
	hex3 := Hex{11 - 321, 33 - 12, 92 - 23}

	if !hex1.Substract(hex2).Equal(hex3) {
		t.Errorf("hex: %+v substract hex: %+v equal %+v", hex1, hex2, hex3)
	}
}
