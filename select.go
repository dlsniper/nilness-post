// This example shows that a call to a method on a pointer value
// that was received via a channel without first checking if the
// value is nil could result to a runtime panic (depending on the
// implementation of the method).

package nilness

import "fmt"

func processUser(ch chan *User) {
	var v *User
	select {
	case v = <-ch:
	default:
	}
	fmt.Printf("got user: %s", v.FullName())
}
