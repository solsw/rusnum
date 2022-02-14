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
