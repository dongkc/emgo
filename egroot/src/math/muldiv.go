package math

func Muldiv(x, m, d uint64) uint64 {
	divx := x / d
	modx := x - divx*d
	divm := m / d
	modm := m - divm*d
	return divx*m + modx*divm + modx*modm/d
}
