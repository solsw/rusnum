package rusnum

import (
	"strconv"
	"strings"
)

// Fraction of a whole number.
type Fraction int

// Supported fractions.
const (
	NoFraction Fraction = iota
	Tenth
	Hundredth
	Thousandth
	Tenthousandth
	Hundredthousandth
	Millionth
	Tenmillionth
	Hundredmillionth
	Milliardth
	Tenmilliardth
	Hundredmilliardth
)

var (
	fracData = [12]struct {
		imultiplier int64
		multiplier  float64
		numberCase  [3]string
	}{
		{1, 1, [3]string{"", "", ""}},
		{10, 10, [3]string{"десятая", "десятых", "десятых"}},
		{100, 100, [3]string{"сотая", "сотых", "сотых"}},
		{1000, 1000, [3]string{"тысячная", "тысячных", "тысячных"}},
		{10000, 10000, [3]string{"десятитысячная", "десятитысячных", "десятитысячных"}},
		{100000, 100000, [3]string{"стотысячная", "стотысячных", "стотысячных"}},
		{1000000, 1000000, [3]string{"миллионная", "миллионных", "миллионных"}},
		{10000000, 10000000, [3]string{"десятимиллионная", "десятимиллионных", "десятимиллионных"}},
		{100000000, 100000000, [3]string{"стомиллионная", "стомиллионных", "стомиллионных"}},
		{1000000000, 1000000000, [3]string{"миллиардная", "миллиардных", "миллиардных"}},
		{10000000000, 10000000000, [3]string{"десятимиллиардная", "десятимиллиардных", "десятимиллиардных"}},
		{100000000000, 100000000000, [3]string{"стомиллиардная", "стомиллиардных", "стомиллиардных"}},
	}
)

func fractionFromStr(sfl string) Fraction {
	ss := strings.Split(sfl, ".")
	if len(ss) != 2 {
		return NoFraction
	}
	if len(ss[1]) > 10 {
		return Hundredmilliardth
	}
	return Fraction(len(ss[1]))
}

func fractionFromFloat(fl float64) Fraction {
	return fractionFromStr(strconv.FormatFloat(fl, 'f', -1, 64))
}

// yields wrong result for 1234567.89
// func fractionFromFloat2(fl float64) Fraction {
// 	_, fr := math.Modf(fl)
// 	if fr == 0 {
// 		return NoFraction
// 	}
// 	f := fl
// 	var fraction Fraction
// 	for {
// 		fraction++
// 		if fraction > Hundredmilliardth {
// 			return Hundredmilliardth
// 		}
// 		f *= 10
// 		_, fr := math.Modf(f)
// 		if fr == 0 {
// 			return fraction
// 		}
// 	}
// }
