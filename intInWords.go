package rusnum

import (
	"math"

	"github.com/solsw/mathhelper"
	"github.com/solsw/slicehelper"
	"github.com/solsw/stringhelper"
)

const (
	zeroInWords     string = "ноль"
	minInt64InWords string = // -9 223 372 036 854 775 808
	"минус девять квинтиллионов двести двадцать три квадриллиона триста семьдесят два триллиона " +
		"тридцать шесть миллиардов восемьсот пятьдесят четыре миллиона семьсот семьдесят пять тысяч восемьсот восемь"
)

// IntInWords returns 'n' in russian words.
//
// If 'withZeros' is false, zero triples are omitted.
// 'gender' determines russian grammatical gender for ones of numbers ending in 1 or 2.
func IntInWords(n int64, withZeros bool, gender GrammaticalGender) string {
	if n == 0 {
		return zeroInWords
	}
	// since -math.MinInt64 is out of int64 range
	if n == math.MinInt64 {
		return minInt64InWords
	}
	absN := mathhelper.Abs(n)
	tt := triplesInWords(absN, withZeros, gender)
	res := stringhelper.JoinSkip(slicehelper.Reverse(tt), " ", func(s string) bool { return s == "" })
	if n < 0 {
		res = "минус " + res
	}
	return res
}
