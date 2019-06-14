// This example shows that if a value and an error are returned
// at the same time and the error is not handled, then the
// could be broken when using the value.

package nilness

func processWindowsLocations() {
	left, err := computePointPosition("left")
	if err != nil {
		panic("err is not nil")
	}

	right, err := computePointPosition("right")
	if left.TopLeft != right.TopLeft {
		panic("windows positions mismatch")
	}
}
