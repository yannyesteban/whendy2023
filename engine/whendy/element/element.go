package element

import (
	"fmt"
	"reflect"
)

var elements = map[string]reflect.Type{}

func Register(name string, e interface{}) {

	a := reflect.TypeOf(e)
	elements[name] = a
}

func New(name string) (interface{}, error) {

	t, ok := elements[name]

	if !ok {
		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}

	i := reflect.New(t).Interface()
	//e := i.(Element)

	return i, nil

}
