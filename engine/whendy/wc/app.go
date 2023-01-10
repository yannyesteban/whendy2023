package wc

import (
	"fmt"

	"sevian.com/whendy/tool"
)

type App struct {
	Method   string
	response []interface{}
}

func (a *App) Init(config map[string]interface{}) {
	tool.ConfigStructure(a, config)
}

func (a *App) EvalMethod() {
	fmt.Println("Que locura!!!!")
	switch a.Method {
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

func (a *App) load() {
	fmt.Println("App Load")

}
