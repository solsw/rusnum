package rusnum

import (
	"fmt"
	"strings"
)

func tripleInWords(triple int64, gender GrammaticalGender) (string, error) {
	if !(0 <= triple && triple <= 999) {
		return "", fmt.Errorf("triple out of range: '%d'", triple)
	}
	if triple == 0 {
		return zeroInWords, nil
	}
	ones := triple % 10           // number of ones
	tens := triple / 10 % 10      // number of tens
	hundreds := triple / 100 % 10 // number of hundreds
	var o string
	if ones > 0 {
		if tens == 1 {
			switch ones {
			case 1:
				o = "одиннадцать"
			case 2:
				o = "двенадцать"
			case 3:
				o = "тринадцать"
			case 4:
				o = "четырнадцать"
			case 5:
				o = "пятнадцать"
			case 6:
				o = "шестнадцать"
			case 7:
				o = "семнадцать"
			case 8:
				o = "восемнадцать"
			case 9:
				o = "девятнадцать"
			}
		} else { // tens != 1
			switch ones {
			case 1:
				switch gender {
				case Neuter:
					o = "одно"
				case Masculine:
					o = "один"
				case Feminine:
					o = "одна"
				}
			case 2:
				switch gender {
				case Feminine:
					o = "две"
				default:
					o = "два"
				}
			case 3:
				o = "три"
			case 4:
				o = "четыре"
			case 5:
				o = "пять"
			case 6:
				o = "шесть"
			case 7:
				o = "семь"
			case 8:
				o = "восемь"
			case 9:
				o = "девять"
			}
		}
	} // if ones > 0
	var t string
	if (tens > 1) || ((tens == 1) && (ones == 0)) {
		switch tens {
		case 1:
			t = "десять"
		case 2:
			t = "двадцать"
		case 3:
			t = "тридцать"
		case 4:
			t = "сорок"
		case 5:
			t = "пятьдесят"
		case 6:
			t = "шестьдесят"
		case 7:
			t = "семьдесят"
		case 8:
			t = "восемьдесят"
		case 9:
			t = "девянoсто"
		}
	} // if (tens > 1) || ((tens == 1) && (ones == 0))
	var h string
	if hundreds > 0 {
		switch hundreds {
		case 1:
			h = "сто"
		case 2:
			h = "двести"
		case 3:
			h = "триста"
		case 4:
			h = "четыреста"
		case 5:
			h = "пятьсот"
		case 6:
			h = "шестьсот"
		case 7:
			h = "семьсот"
		case 8:
			h = "восемьсот"
		case 9:
			h = "девятьсот"
		}
	} // if hundreds > 0
	var sb strings.Builder
	if len(h) > 0 {
		sb.WriteString(h)
	}
	if len(t) > 0 {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(t)
	}
	if len(o) > 0 {
		if sb.Len() > 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(o)
	}
	return sb.String(), nil
}

// triplesInWords returns triples comprising 'n' in russian words.
func triplesInWords(n int64, withZeros bool, gender GrammaticalGender) []string {
	// 'n' must not be negative
	if n == 0 {
		return []string{zeroInWords}
	}
	var tt []string
	tripleNumber := 0
	for n > 0 {
		triple := n % 1000
		t := ""
		if triple > 0 || withZeros {
			if tripleNumber == 0 {
				t, _ = tripleInWords(triple, gender)
			} else {
				if tripleNumber == 1 {
					t, _ = tripleInWords(triple, Feminine)
				} else {
					t, _ = tripleInWords(triple, Masculine)
				}
				nnc := getNumeralNumberCase(triple)
				switch tripleNumber {
				case 1:
					t += " " + thousandNumberCase[nnc]
				case 2:
					t += " " + millionNumberCase[nnc]
				case 3:
					t += " " + milliardNumberCase[nnc]
				case 4:
					t += " " + trillionNumberCase[nnc]
				case 5:
					t += " " + quadrillionNumberCase[nnc]
				case 6:
					t += " " + quintillionNumberCase[nnc]
				}
			}
		}
		tt = append(tt, t)
		tripleNumber++
		n /= 1000
	}
	return tt
}
