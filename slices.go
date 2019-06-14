// This example shows that when accessing an index of a slice
// that has not been initialized, it will lead to a runtime error.

package nilness

import "fmt"

func sliceIndex() {
	var s []int
	fmt.Printf("first value: %d", s[0])
}
