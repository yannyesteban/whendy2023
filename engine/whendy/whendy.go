package whendy

import (
	"fmt"
	"net/http"
	"reflect"

	"sevian.com/whendy/element"
	"sevian.com/whendy/tool"
)

var scheme = map[string]reflect.Type{}

type InfoElement struct {
	Element  string
	Id       string
	Eparams  map[string]interface{}
	AppendTo string
	SetPanel string
	Method   string
	Name     string
}

func (i *InfoElement) Set(m map[string]interface{}) {
	var value interface{}
	var ok bool

	fmt.Println(" es element: ", m["element"])
	if value, ok = m["element"]; ok {
		i.Element = value.(string)
	}

	if value, ok = m["id"]; ok {
		i.Id = value.(string)
	}

	if value, ok = m["name"]; ok {
		i.Name = value.(string)
	}

	if value, ok = m["method"]; ok {
		i.Method = value.(string)
	}
	//i.Id = m["id"].(string)
	i.Eparams = m["eparams"].(map[string]interface{})

}

type Whendy struct {
	Id     string
	Name   string
	Method string
	Edad   int

	w   http.ResponseWriter
	req *http.Request

	Ini map[string]interface{}

	response []interface{}
}

func (w *Whendy) Init(config map[string]interface{}) {

	tool.ConfigStructure(w, config)

}

func (w *Whendy) Element(e Element) {

	e.EvalMethod()
}

func (Whendy) Render1() string {

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

func (wh Whendy) EvalStartMode() {

	mode := wh.req.Header.Get("Application-Mode")

	if mode == "start" {
		fmt.Println(wh.req.Header.Get("Application-Mode"))
		info := InfoElement{}
		info.Set(wh.Ini)
		wh.setElement(info)
	}

}
func (whendy Whendy) Render() {

	whendy.EvalStartMode()
	w := whendy.w

	element.Register("unit", Unit{})

	k, _ := element.New("unit2")

	ee := k.(Element)
	ee.EvalMethod()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Request-Method, Application-Mode, authorization, sid,  Application-Id")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Allow", "GET, POST, OPTIONS, PUT, DELETE")

	w.Header().Set("Content-Type", "application/json")

	whendy.getResponse()
	fmt.Fprintf(w, `[9,4,8,16,32]`)

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
	whendy.w = w
	whendy.req = req

	fmt.Println(req.Header.Get("Application-Mode"))

	//q, _ := req.Cookie("accessToken")
	//println(q.Value)

	//sess.Set("username", "yanny1")
	//username := sess.Get("username")

	//fmt.Println(username)
	//st.LoadJson("configuration/constants.json")

}

func (whendy *Whendy) setElement(info InfoElement) {

	fuel := tool.LoadJsonFile(info.Name)

	fmt.Println("...", fuel["templateFile"])
	typ, _ := element.New(info.Element)

	ele := typ.(Element)
	ele.Init(fuel)
	ele.EvalMethod()

	fmt.Println(" QHE INFO ", info.Id)
	//int64(info.Eparams["cedula"].(float64)
}

func (whendy *Whendy) getResponse() []interface{} {
	return whendy.response
}
