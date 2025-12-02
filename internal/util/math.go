package util

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Wrap(value, max int) (wrapped, wraps int) {
	wrapped = ((value % max) + max) % max

	if value >= 0 {
		wraps = value / max
	} else {
		// How many complete cycles of max do we go through?
		// -1 is 0 wraps (stays in range after wrapping)
		// -100 is 1 wrap, -101 is 1 wrap, -200 is 2 wraps
		wraps = (-value-1)/max + 1
	}

	return wrapped, wraps
}

// WrapValue is a convenience function if you only need the wrapped value
func WrapValue(value, max int) int {
	wrapped, _ := Wrap(value, max)
	return wrapped
}

// WrapRange wraps a value to stay within [min, max]
// Examples: WrapRange(10, 0, 9) = 0, WrapRange(-1, 0, 9) = 9
func WrapRange(value, min, max int) int {
	rangeSize := max - min + 1
	value = ((value - min) % rangeSize)
	if value < 0 {
		value += rangeSize
	}
	return value + min
}
