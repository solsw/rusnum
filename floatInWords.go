package rusnum

import (
	"math"
	"strings"
)

// Binder of whole and fractional parts.
type (
	Binder int
)

// Supported binders.
const (
	_ Binder = iota
	And
	Whole
)

var (
	wholeNumberCase = [3]string{"целая", "целых", "целых"}
)

// FloatInWords returns 'f' represented in russian words.
// If f's integer part is 0 and 'zeroInt' is false, no integer part is returned.
// If f's fractional part is 0 and 'zeroFrac' is false, no fractional part is returned (see FracInWords).
// 'withZero' is used by f's integer part (see IntInWords).
func FloatInWords(f float64, fraction Fraction, binder Binder, zeroInt, zeroFrac, withZero bool) string {
	absF := math.Abs(f)
	fint, frac := math.Modf(absF)
	ifint := int64(fint)
	sfrac := FracInWords(frac, fraction, zeroFrac)
	var sb strings.Builder
	if ifint == 0 && !zeroInt && len(sfrac) > 0 {
		sb.WriteString(sfrac)
	} else {
		gender := Masculine
		if binder == Whole {
			gender = Feminine
		}
		sb.WriteString(IntInWords(ifint, withZero, gender))
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
