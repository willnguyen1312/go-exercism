package triangle

import (
	"math"
)

// KindFromSides function
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}

	for _, side := range sides {
		if side <= 0 || math.IsNaN(side) || math.IsInf(side, 0) {
			return NaT
		}
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

// Kind type
type Kind int

const (
	// NaT - not a triangle
	NaT = iota
	// Equ - equilateral
	Equ
	// Iso - isosceles
	Iso
	// Sca - scalene
	Sca
)
