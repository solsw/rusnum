package rusnum

import (
	"fmt"
)

type (
	// GrammaticalGender - russian grammatical gender.
	GrammaticalGender int
)

const (
	// Neuter gender.
	Neuter GrammaticalGender = iota
	// Masculine gender.
	Masculine
	// Feminine gender.
	Feminine
)

type (
	// russian number/case used with numbers
	// русские число/падеж, используемые с числами
	numeralNumberCase int
)

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
	// depends on two last digits in general
	// в общем случае определяется двумя последними цифрами
	if last2 == 11 || last2 == 12 || last2 == 13 || last2 == 14 {
		return pluralGenitive
	}
	// depends on one last digit now
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
	if absN < 0 {
		absN = -absN
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
// If 'n' is 0 and 'withZeros' is false, empty string is returned.
// If 'withZeros' is false, zero triples are omitted.
func IntInWordsAndItems(n int64, showZero, withZeros bool, gender GrammaticalGender, items string) string {
	if n == 0 && !showZero {
		return ""
	}
	return IntInWords(n, withZeros, gender) + " " + items
}
