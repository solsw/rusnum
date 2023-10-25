package rusnum

import (
	"math"
	"strings"
)

// Binder of integer and fractional parts.
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
// If f's fractional part is 0 and 'zeroFrac' is false, no fractional part is returned (see [FloatInFractions]).
// 'withZeros' is used by f's integer part (see [IntInWords]).
func FloatInWords(f float64, frac Fraction, binder Binder, zeroInt, zeroFrac, withZeros bool) string {
	if frac < NoFraction || frac > Tenmilliardth {
		frac = Hundredmilliardth
	}
	fracDatum := fracData[frac]
	absF := math.Abs(f)
	wholesF := int64(math.Floor(absF))
	numOfFracs := int64(math.Floor(absF*fracDatum.multiplier)) - wholesF*fracDatum.imultiplier
	sfrac := ""
	if frac != NoFraction {
		sfrac = IntInWords(numOfFracs, false, Feminine) + " " + fracDatum.numberCase[getNumeralNumberCase(numOfFracs)]
	}
	var sb strings.Builder
	if wholesF == 0 && !zeroInt && len(sfrac) > 0 {
		sb.WriteString(sfrac)
	} else {
		gender := Masculine
		if binder == Whole {
			gender = Feminine
		}
		sb.WriteString(IntInWords(wholesF, withZeros, gender))
		switch binder {
		case And:
			if len(sfrac) > 0 {
				sb.WriteString(" и")
			}
		case Whole:
			sb.WriteString(" ")
			sb.WriteString(wholeNumberCase[getNumeralNumberCase(wholesF)])
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

// FloatInWordsAuto is like [FloatInWords] but determines the [Fraction] automatically.
func FloatInWordsAuto(f float64, binder Binder, zeroInt, zeroFrac, withZeros bool) string {
	return FloatInWords(f, fractionFromFloat(f), binder, zeroInt, zeroFrac, withZeros)
}

// FloatInFractions returns 'f' expressed in 'fracs' in russian words.
// If result is 0 and 'showZero' is false, empty string is returned.
func FloatInFractions(f float64, frac Fraction, showZero bool) string {
	if frac == NoFraction {
		return IntInWords(int64(f), showZero, Masculine)
	}
	if frac < NoFraction || frac > Tenmilliardth {
		frac = Hundredmilliardth
	}
	fracDatum := fracData[frac]
	numOfFracs := int64(math.Floor(math.Abs(f) * fracDatum.multiplier))
	if numOfFracs == 0 && !showZero {
		return ""
	}
	r := IntInWords(numOfFracs, false, Feminine) + " " + fracDatum.numberCase[getNumeralNumberCase(numOfFracs)]
	if f < 0 {
		return "минус " + r
	}
	return r
}

// FloatInFractionsAuto is like [FloatInFractions] but determines the [Fraction] automatically.
func FloatInFractionsAuto(f float64, showZero bool) string {
	return FloatInFractions(f, fractionFromFloat(f), showZero)
}
