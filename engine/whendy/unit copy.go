package whendy

import (
	"fmt"

	"sevian.com/tool"
)

type Device struct {
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

func (obj *Device) Init(config map[string]interface{}) {

	tool.ConfigStructure(obj, config)

}

func (Device) EvalMethod() {
	fmt.Println("Device App")
}
