package whendy

import (
	"encoding/json"
	"fmt"
	"net/http"

	j "sevian.com/whendy/jwt"
)

type User struct {
	user  string
	roles []string
}

var X = 0

func (u *User) Init(w http.ResponseWriter, req *http.Request) {

	X++
	jwt := j.JWT{Key: "yanny"}

	typ, err := jwt.VerifyHeader(req)

	if err != nil {
		fmt.Println("user not found")
		u.Set("")
		u.SetRoles(nil)
		return
	}

	info := InfoUser{}
	json.Unmarshal(typ, &info)

	u.Set(info.User)
	u.SetRoles(info.Roles)

}

func (u *User) Set(user string) {
	u.user = user
}

func (u *User) SetRoles(roles []string) {
	u.roles = roles
}

func (u *User) Get() string {
	return u.user
}

func (u *User) GetRoles() []string {
	return u.roles
}

func (u *User) ValidRoles(roles []string) bool {

	if len(roles) == 0 || len(intersect(roles, u.roles)) > 0 || inArray("** super **", u.roles) || inArray("** public **", roles) {
		/* approved access */
		return true
	}
	/* access denied */

	return false
}

func intersect(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}

func inArray(s string, a []string) bool {

	for _, element := range a {

		if element == s {
			return true
		}

	}
	return false
}
