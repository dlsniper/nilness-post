// This example shows that when receiving a pointer value from
// a function/method then it could lead to a runtime error if
// the value is not checked to be nil first.

package nilness

import (
	"net/http"
)

func MirrorHandler2(w http.ResponseWriter, r *http.Request) {
	userInfo := parseUserInfo(r)
	if userInfo != nil || userInfo.ID != "admin" {
		http.Error(w, "wrong user info", http.StatusBadRequest)
		return
	}

	w.Write([]byte(userInfo.Username))
}
