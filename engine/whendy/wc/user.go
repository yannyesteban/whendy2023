package wc

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"sevian.com/whendy"
	"sevian.com/whendy/tool"
)

type User struct {
	Id           string
	Name         string
	TemplateFile string
	ClassName    string
	Elements     []whendy.InfoElement

	CssSheets []string
	JsModules []string
	Modules   []JsModule

	Method   string
	response []interface{}

	store *whendy.Store
	info  whendy.InfoUser
	error int

	//ConfigStructure func(interface{}, map[string]interface{})
}

func (a *User) Init(info whendy.InfoElement) {

	str := tool.LoadFile(info.Name)

	json.Unmarshal([]byte(str), &a)

}

func (a *User) EvalMethod(method string) {

	fmt.Println("SEsion ", a.store.GetSes("APP_PATH"))
	switch method {
	case "login":
		a.login()

	case "load":
		a.load()
	}

}

func (a *User) GetResponse() []interface{} {

	return a.response
}

func (a *User) AddResponse(response interface{}) {
	a.response = append(a.response, response)
}

func (a *User) login() {
	db := a.store.GetDB("gt") //sql.Open("mysql", "root:123456@/gt")

	//table := "_sg_users"
	rUser := a.store.GetReq("user").(string)
	rPass := a.store.GetReq("password").(string)
	/*
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	*/
	//defer db.Close()

	stmtOut, err := db.Prepare("SELECT u.user, u.pass FROM _sg_users as u WHERE u.user=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var user string
	var pass string

	err = stmtOut.QueryRow(rUser).Scan(&user, &pass) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	if md5Hash(rPass) == pass {

		a.error = 0

		a.info.User = user
		a.info.Roles = a.getRoles(user)
	} else {
		a.info.User = ""
		a.info.Roles = nil
		a.error = 1
	}

}

func (a *User) getRoles(user string) []string {

	db := a.store.GetDB("gt") //sql.Open("mysql", "root:123456@/gt")

	//table := "_sg_grp_usr"

	var roles []string

	//defer db.Close()

	stmtOut, err := db.Prepare("SELECT g.group FROM _sg_grp_usr as g WHERE user=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	res, err2 := stmtOut.Query(user)
	//defer res.Close()

	if err2 != nil {
		log.Fatal(err2)
	}
	var role string
	for res.Next() {

		err = res.Scan(&role)

		if err != nil {
			log.Fatal(err)
		}

		roles = append(roles, role)
		fmt.Printf("%v\n", roles)
	}

	return roles

}

func (app *User) load() {

	template := tool.LoadFile(app.TemplateFile)

	data := map[string]interface{}{

		"mode":    "update",
		"element": "wh-app",
		"id":      app.Id,
		"props": map[string]interface{}{
			"name":      app.Name,
			"element":   "wh-app",
			"className": app.ClassName,
			"modules":   app.Modules,
			"jsModules": app.JsModules,
			"innerHTML": template,
		},
	}

	app.AddResponse(data)

}

func (app *User) GetElements() []whendy.InfoElement {

	return app.Elements

}

func (app *User) SetStore(store *whendy.Store) {
	app.store = store
}

func (app *User) GetUserInfo() whendy.InfoUser {

	return app.info
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
