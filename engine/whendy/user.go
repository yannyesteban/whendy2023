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
	Key   string
}

func (u *User) Init(w http.ResponseWriter, req *http.Request) {

	jwt := j.JWT{Key: u.Key}

	typ, err := jwt.VerifyHeader(req)

	if err != nil {
		fmt.Println("user not found")
		u.user = ""
		u.roles = nil
		return
	}

	info := InfoUser{}
	json.Unmarshal(typ, &info)

	u.user = info.User
	u.roles = info.Roles

}

func (u *User) Set(info InfoUser) string {
	u.user = info.User
	u.roles = info.Roles

	jwt := j.JWT{Key: u.Key}

	return jwt.Generate(info)

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
