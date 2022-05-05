package rusnum

import (
	"fmt"
)

func ExampleIntInWordsAndItems() {
	fmt.Println(IntInWordsAndItems(2, false, false, Masculine, "кирпича"))
	// Output:
	// два кирпича
}
