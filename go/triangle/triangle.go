package triangle

// Kind type
type Kind int

const (
	// NaT value
	NaT = iota
	// Equ value
	Equ
	// Iso value
	Iso
	// Sca value
	Sca
)

// KindFromSides function to determine triangle type based on three sides
func KindFromSides(a, b, c float64) Kind {
	if a+b >= c || a+c >= b || b+c >= a {
		return NaT
	}

	if a == b && a == c && b == c {
		return Equ
	}

	if (a == b && c == b) || (a == c && b == c) || (b == a && c == a) {
		return Iso
	}

	return Sca
}
