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

// FloatInWords returns 'f' in russian words.
// If f's integer part is 0 and 'zeroInt' is false, no integer part is returned.
// If f's fractional part is 0 and 'zeroFrac' is false, no fractional part is returned (see FloatInFractions).
// 'withZeros' is used by f's integer part (see IntInWords).
func FloatInWords(f float64, frac Fraction, binder Binder, zeroInt, zeroFrac, withZeros bool) string {
	if frac < NoFraction || frac > Tenmilliardth {
		frac = Hundredmilliardth
	}
	fracDatum := fracData[frac]
	absF := math.Abs(f)
	fint := math.Floor(absF)
	numOfFracs := math.Floor(absF*fracDatum.multiplier - fint*fracDatum.multiplier)
	inumOfFracs := int64(numOfFracs)
	sfrac := ""
	if frac != NoFraction {
		sfrac = IntInWords(inumOfFracs, false, Feminine) + " " + fracDatum.numberCase[getNumeralNumberCase(inumOfFracs)]
	}
	var sb strings.Builder
	ifint := int64(fint)
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
