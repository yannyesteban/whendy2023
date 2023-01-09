package whendy

import (
	"fmt"
	"net/http"
	"reflect"

	"sevian.com/tool"
)

var scheme = map[string]reflect.Type{}

type Whendy struct {
	Id     string
	Name   string
	Method string
	Edad   int

	w   http.ResponseWriter
	req *http.Request
}

func (w *Whendy) Init(config map[string]interface{}) {

	tool.ConfigStructure(w, config)

}

func (w *Whendy) Element(e Element) {

	e.EvalMethod()
}

func (Whendy) Render() string {

	RegisterType("device", Device{})
	RegisterType("unit", Unit{})

	el, err := NewElement("unit")

	if err == nil {
		el.EvalMethod()
	} else {
		fmt.Println(err)
	}

	return "render HTML"
}

func New(i interface{}) (interface{}, error) {

	t := reflect.TypeOf(i)

	return reflect.New(t).Interface(), nil
}

func RegisterType(name string, t interface{}) {
	a := reflect.TypeOf(t)
	scheme[name] = a
}

func NewElement(name string) (Element, error) {

	t, ok := scheme[name]

	if !ok {
		return nil, fmt.Errorf("unrecognized type name: %s", name)
	}

	i := reflect.New(t).Interface()
	e := i.(Element)

	return e, nil

}

func (whendy *Whendy) Start(w http.ResponseWriter, req *http.Request) {

	fmt.Println(req.Header.Get("Application-Mode"))

	//q, _ := req.Cookie("accessToken")
	//println(q.Value)

	//sess.Set("username", "yanny1")
	//username := sess.Get("username")

	//fmt.Println(username)
	//st.LoadJson("configuration/constants.json")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Request-Method, Application-Mode, authorization, sid,  Application-Id")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Allow", "GET, POST, OPTIONS, PUT, DELETE")

	w.Header().Set("Content-Type", "application/json")

	//defer sess.SessionRelease(w)

	fmt.Fprintf(w, `[6,7,8]`)

}
