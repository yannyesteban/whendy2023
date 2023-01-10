package whendy

import (
	"fmt"

	"sevian.com/whendy/tool"
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

func (obj *Unit) Init(config map[string]interface{}) {

	tool.ConfigStructure(obj, config)

}

func (Unit) EvalMethod() {
	fmt.Println("Unit App 2023")
}
