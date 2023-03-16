package rusnum

import (
	"github.com/solsw/mathhelper"
)

// OnesInWords returns ones (three lower digits) of 'n' in russian words.
//
// 'gender' determines russian grammatical gender for ones of numbers ending in 1 or 2.
func OnesInWords(n int64, gender GrammaticalGender) string {
	absN := mathhelper.Abs(n)
	tt := triplesInWords(absN, true, gender)
	return tt[0]
}

func notOnesInWords(n, min, tripleNum int64) string {
	absN := mathhelper.Abs(n)
	if absN < min {
		return ""
	}
	tt := triplesInWords(absN, true, Neuter)
	return tt[tripleNum]
}

// ThousandsInWords returns thousands (next three digits after ones) of 'n' in russian words.
func ThousandsInWords(n int64) string {
	return notOnesInWords(n, 1000, 1)
}

// MillionsInWords returns millions (next three digits after thousands) of 'n' in russian words.
func MillionsInWords(n int64) string {
	return notOnesInWords(n, 1000000, 2)
}

// MilliardsInWords returns milliards (next three digits after millions) of 'n' in russian words.
func MilliardsInWords(n int64) string {
	return notOnesInWords(n, 1000000000, 3)
}

// TrillionsInWords returns trillions (next three digits after milliards) of 'n' in russian words.
func TrillionsInWords(n int64) string {
	return notOnesInWords(n, 1000000000000, 4)
}

// QuadrillionsInWords returns quadrillions (next three digits after trillions) of 'n' in russian words.
func QuadrillionsInWords(n int64) string {
	return notOnesInWords(n, 1000000000000000, 5)
}

// QuintillionsInWords returns quintillions (next three digits after quadrillions) of 'n' in russian words.
func QuintillionsInWords(n int64) string {
	return notOnesInWords(n, 1000000000000000000, 6)
}
