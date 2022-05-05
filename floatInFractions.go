package rusnum

import (
	"math"
)

var (
	fracData = [12]struct {
		multiplier float64
		numberCase [3]string
	}{
		{1, [3]string{"", "", ""}},
		{10, [3]string{"десятая", "десятых", "десятых"}},
		{100, [3]string{"сотая", "сотых", "сотых"}},
		{1000, [3]string{"тысячная", "тысячных", "тысячных"}},
		{10000, [3]string{"десятитысячная", "десятитысячных", "десятитысячных"}},
		{100000, [3]string{"стотысячная", "стотысячных", "стотысячных"}},
		{1000000, [3]string{"миллионная", "миллионных", "миллионных"}},
		{10000000, [3]string{"десятимиллионная", "десятимиллионных", "десятимиллионных"}},
		{100000000, [3]string{"стомиллионная", "стомиллионных", "стомиллионных"}},
		{1000000000, [3]string{"миллиардная", "миллиардных", "миллиардных"}},
		{10000000000, [3]string{"десятимиллиардная", "десятимиллиардных", "десятимиллиардных"}},
		{100000000000, [3]string{"стомиллиардная", "стомиллиардных", "стомиллиардных"}},
	}
)

// FloatInFractions returns 'f' expressed in 'frac's in russian words.
// If result is 0 and 'showZero' is false, empty string is returned.
func FloatInFractions(f float64, frac Fraction, showZero bool) string {
	if frac == NoFraction {
		return IntInWords(int64(f), showZero, Masculine)
	}
	if frac < NoFraction || frac > Tenmilliardth {
		frac = Hundredmilliardth
	}
	fracDatum := fracData[frac]
	numOfFracs := math.Floor(math.Abs(f) * fracDatum.multiplier)
	inumOfFracs := int64(numOfFracs)
	if inumOfFracs == 0 && !showZero {
		return ""
	}
	r := IntInWords(inumOfFracs, false, Feminine) + " " + fracDatum.numberCase[getNumeralNumberCase(inumOfFracs)]
	if f < 0 {
		return "минус " + r
	}
	return r
}

// FloatInFractionsAuto is like FloatInFractions but determines the Fraction automatically.
func FloatInFractionsAuto(f float64, showZero bool) string {
	return FloatInFractions(f, fractionFromFloat(f), showZero)
}
