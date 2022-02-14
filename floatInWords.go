package rusnum

import (
	"math"
	"strings"
)

// Binder of whole and fractional parts.
type Binder int

const (
	NoBinder Binder = iota
	// And separates integer and fractional parts with russian for 'and'.
	And
	// Whole separates integer and fractional parts with russian for 'whole'.
	Whole
)

var (
	wholeNumberCase = [3]string{"целая", "целых", "целых"}
)

// FloatInWords returns 'f' represented in russian words.
// If f's integer part is 0 and 'zeroInt' is false, no integer part is returned.
// If f's fractional part is 0 and 'zeroFrac' is false, no fractional part is returned (see FracInWords).
// 'withZeros' is used by f's integer part (see IntInWords).
func FloatInWords(f float64, fraction Fraction, binder Binder, zeroInt, zeroFrac, withZeros bool) string {
	absF := math.Abs(f)
	fint, frac := math.Modf(absF)
	ifint := int64(fint)
	sfrac := ""
	if fraction != NoFraction {
		sfrac = FracInWords(frac, fraction, zeroFrac)
	}
	var sb strings.Builder
	if ifint == 0 && !zeroInt && len(sfrac) > 0 {
		sb.WriteString(sfrac)
	} else {
		gender := Masculine
		if binder == Whole {
			gender = Feminine
		}
		sb.WriteString(IntInWords(ifint, withZeros, gender))
		switch binder {
		case And:
			if len(sfrac) > 0 {
				sb.WriteString(" и")
			}
		case Whole:
			sb.WriteString(" ")
			sb.WriteString(wholeNumberCase[getNumeralNumberCase(ifint)])
		}
		if len(sfrac) > 0 {
			sb.WriteString(" ")
			sb.WriteString(sfrac)
		}
	}
	if f < 0 {
		return "минус " + sb.String()
	}
	return sb.String()
}

// FloatInWordsAuto is like FloatInWords but determines the Fraction automatically.
func FloatInWordsAuto(f float64, binder Binder, zeroInt, zeroFrac, withZeros bool) string {
	return FloatInWords(f, fractionFromFloat(f), binder, zeroInt, zeroFrac, withZeros)
}
