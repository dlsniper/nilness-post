// This file contains all the boilerplate needed in the res of the files.

package nilness

import (
	"encoding/json"
	"net/http"
)

func _() {
	sliceIndex()
	processUser(nil)
	processWindowsLocations()
	MirrorHandler(nil, nil)
	MirrorHandler2(nil, nil)
}

type User struct {
	ID   string
	Name string
}

func (u *User) FullName() string { return "" }

type Point struct {
	TopLeft int
}

func computePointPosition(windowName string) (*Point, error) {
	_ = windowName
	return nil, nil
}

type UserInfo struct {
	ID       string
	Username string
}

func parseUserInfo(r *http.Request) *UserInfo {
	userInfo := &UserInfo{}
	err := json.NewDecoder(r.Body).Decode(userInfo)
	if err != nil {
		panic(err)
	}

	return nil
}
