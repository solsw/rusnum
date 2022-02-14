package rusnum

import (
	"math"
)

var (
	tenthNumberCase             = [3]string{"десятая", "десятых", "десятых"}
	hundredthNumberCase         = [3]string{"сотая", "сотых", "сотых"}
	thousandthNumberCase        = [3]string{"тысячная", "тысячных", "тысячных"}
	tenthousandthNumberCase     = [3]string{"десятитысячная", "десятитысячных", "десятитысячных"}
	hundredthousandthNumberCase = [3]string{"стотысячная", "стотысячных", "стотысячных"}
	millionthNumberCase         = [3]string{"миллионная", "миллионных", "миллионных"}
	tenmillionthNumberCase      = [3]string{"десятимиллионная", "десятимиллионных", "десятимиллионных"}
	hundredmillionthNumberCase  = [3]string{"стомиллионная", "стомиллионных", "стомиллионных"}
	milliardthNumberCase        = [3]string{"миллиардная", "миллиардных", "миллиардных"}
	tenmilliardthNumberCase     = [3]string{"десятимиллиардная", "десятимиллиардных", "десятимиллиардных"}
	hundredmilliardthNumberCase = [3]string{"стомиллиардная", "стомиллиардных", "стомиллиардных"}
)

// FracInWords returns 'frac' expressed in 'fraction's in russian words.
// If result is 0 and 'showZero' is false, empty string is returned.
func FracInWords(frac float64, fraction Fraction, showZero bool) string {
	if fraction == NoFraction {
		return IntInWords(int64(frac), showZero, Masculine)
	}
	absFrac := math.Abs(frac)
	var numberCase [3]string
	switch fraction {
	case Tenth:
		absFrac *= 10
		numberCase = tenthNumberCase
	case Hundredth:
		absFrac *= 100
		numberCase = hundredthNumberCase
	case Thousandth:
		absFrac *= 1000
		numberCase = thousandthNumberCase
	case Tenthousandth:
		absFrac *= 10000
		numberCase = tenthousandthNumberCase
	case Hundredthousandth:
		absFrac *= 100000
		numberCase = hundredthousandthNumberCase
	case Millionth:
		absFrac *= 1000000
		numberCase = millionthNumberCase
	case Tenmillionth:
		absFrac *= 10000000
		numberCase = tenmillionthNumberCase
	case Hundredmillionth:
		absFrac *= 100000000
		numberCase = hundredmillionthNumberCase
	case Milliardth:
		absFrac *= 1000000000
		numberCase = milliardthNumberCase
	case Tenmilliardth:
		absFrac *= 10000000000
		numberCase = tenmilliardthNumberCase
	// all other fractions are treated as Hundredmilliardth
	default:
		absFrac *= 100000000000
		numberCase = hundredmilliardthNumberCase
	}
	fint, _ := math.Modf(absFrac)
	ifint := int64(fint)
	if ifint == 0 && !showZero {
		return ""
	}
	r := IntInWords(ifint, false, Feminine) + " " + numberCase[getNumeralNumberCase(ifint)]
	if frac < 0 {
		return "минус " + r
	}
	return r
}

// FracInWordsAuto is like FracInWords but determines the Fraction automatically.
func FracInWordsAuto(frac float64, showZero bool) string {
	return FracInWords(frac, fractionFromFloat(frac), showZero)
}
