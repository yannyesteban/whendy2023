package whendy

import (
	"fmt"
)

type Unit struct {
	Id          string
	Name        string
	Eparams     map[string]interface{}
	SetPanel    string
	AppendTo    string
	Method      string
	ReplayToken string
	OnDesing    bool
	OnDebug     bool

	Response      []interface{}
	AcceptedRoles []string
}

func (obj *Unit) Init(info InfoElement) {

	//tool.ConfigStructure(obj, config)

}

func (Unit) EvalMethod(method string) {
	fmt.Println("Unit App 2023")
}

func (a *Unit) GetResponse() []interface{} {

	return a.Response
}

func (a *Unit) AddResponse(response interface{}) {
	a.Response = append(a.Response, response)
}
