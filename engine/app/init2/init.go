package init2

import (
	"sevian.com/whendy/element"
	"sevian.com/whendy/wc"
)

func Start() {
	element.Register("user", wc.User{})
	element.Register("app", wc.App{})

}
