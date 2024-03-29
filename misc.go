package rusnum

import (
	"fmt"
)

// GrammaticalGender - russian grammatical gender.
type GrammaticalGender int

const (
	// Neuter grammatical gender.
	Neuter GrammaticalGender = iota
	// Masculine grammatical gender.
	Masculine
	// Feminine grammatical gender.
	Feminine
)

// russian number/case used with numbers
// русские число/падеж, используемые с числами
type numeralNumberCase int

// numeralNumberCase choices
const (
	// singular, nominative case
	// единственное число, именительный падеж (1, 21 час (но 11 часов))
	singularNominative numeralNumberCase = iota
	// singular, genitive case
	// единственное число, родительный падеж (2, 3, 4, 22 часа (но 12, 13, 14 часов))
	singularGenitive
	// plural, genitive case
	// множественное число, родительный падеж (0, 5 (и всё остальное) часов)
	pluralGenitive
)

func getNumeralNumberCasePrim(last2 int64) numeralNumberCase {
	// 0 <= last2 <= 99
	// in general case depends on two last digits
	// в общем случае определяется двумя последними цифрами
	if last2 == 11 || last2 == 12 || last2 == 13 || last2 == 14 {
		return pluralGenitive
	}
	// now depends on one last digit
	// теперь определяется одной последней цифрой
	switch last2 % 10 {
	case 1:
		return singularNominative
	case 2, 3, 4:
		return singularGenitive
	default:
		return pluralGenitive
	}
}

func getNumeralNumberCase(n int64) numeralNumberCase {
	absN := n
	if n < 0 {
		absN = -n
	}
	return getNumeralNumberCasePrim(absN % 100)
}

// IntAndItems returns string containing 'n' and 'items'.
// If 'n' is 0 and 'showZero' is false, empty string is returned.
func IntAndItems(n int64, showZero bool, items string) string {
	if n == 0 && !showZero {
		return ""
	}
	return fmt.Sprintf("%d %s", n, items)
}

// IntInWordsAndItems returns string containing 'n' in russian words and 'items'.
// If 'n' is 0 and 'showZero' is false, empty string is returned.
// If 'withZeros' is false, zero triples are omitted.
func IntInWordsAndItems(n int64, showZero, withZeros bool, gender GrammaticalGender, items string) string {
	if n == 0 && !showZero {
		return ""
	}
	return IntInWords(n, withZeros, gender) + " " + items
}
