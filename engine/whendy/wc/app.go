package wc

import (
	"encoding/json"

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

	//ConfigStructure func(interface{}, map[string]interface{})
}

func (a *App) Init(info whendy.InfoElement) {

	str := tool.LoadFile(info.Name)

	json.Unmarshal([]byte(str), &a)

}

func (a *App) EvalMethod(method string) {

	switch method {
	case "init":
		a.load()
	}
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
