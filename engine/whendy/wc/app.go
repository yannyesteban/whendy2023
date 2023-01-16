package wc

import (
	"encoding/json"
	"fmt"

	"sevian.com/whendy"
	"sevian.com/whendy/tool"
)

type JsModule struct {
	Src       string `json:"src"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Component string `json:"component"`
}
type App struct {
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

	//ConfigStructure func(interface{}, map[string]interface{})
}

func (a *App) Init(info whendy.InfoElement) {

	str := tool.LoadFile(info.Name)

	json.Unmarshal([]byte(str), &a)

}

func (a *App) EvalMethod(method string) {

	fmt.Println("SEsion ", a.store.GetSes("APP_PATH"))
	switch method {
	case "init":
		a.load()
	}

	db := a.store.GetDB("gt") //sql.Open("mysql", "root:123456@/gt")
	/*
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	*/
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id, name  FROM unit u where id=?")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	var unitId int
	var unitName string // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(2023).Scan(&unitId, &unitName) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", unitId)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(2024).Scan(&unitId, &unitName) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 2024 is: %v %v\n", unitId, unitName)
}

func (a *App) GetResponse() []interface{} {

	return a.response
}

func (a *App) AddResponse(response interface{}) {
	a.response = append(a.response, response)
}

func (app *App) load() {

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

func (app *App) GetElements() []whendy.InfoElement {

	return app.Elements

}

func (app *App) SetStore(store *whendy.Store) {
	app.store = store
}
