// This example shows that the "userInfo" is always not nil
// since it's initialized and then it can't be made nil anymore.

package nilness

import (
	"encoding/json"
	"net/http"
)

func MirrorHandler(w http.ResponseWriter, r *http.Request) {
	userInfo := &UserInfo{}
	err := json.NewDecoder(r.Body).Decode(userInfo)
	if err != nil {
		http.Error(w, "Encountered error", http.StatusBadRequest)
		return
	}

	if userInfo == nil {
		http.Error(w, "nil user info", http.StatusBadRequest)
		return
	}

	w.Write([]byte(userInfo.Username))
}
