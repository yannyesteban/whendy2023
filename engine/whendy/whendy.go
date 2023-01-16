package whendy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"sevian.com/whendy/element"
	"sevian.com/whendy/tool"
)

type Element interface {
	//LoadConfig(string)
	EvalMethod(method string)
	Init(info InfoElement)
	GetResponse() []interface{}
	AddResponse(response interface{})
	SetStore(*Store)
}

var scheme = map[string]reflect.Type{}

type InfoUser struct {
	User  string
	Roles []string
}

type AppComponent interface {
	AppendTo(string)
	SetPanel(string)
	SetParams(map[string]interface{})
	SetMethod(string)
}

type UserAdmin interface {
	GetUserInfo() InfoUser
}

type ElementAdmin interface {
	GetElements() []InfoElement
}

type InfoElement struct {
	Element  string
	Id       string
	Eparams  map[string]interface{}
	AppendTo string
	SetPanel string
	Method   string
	Name     string
}

/*
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
*/
type Whendy struct {
	IniSource string
	Id        string
	Name      string
	Method    string

	w   http.ResponseWriter
	req *http.Request

	Ini map[string]interface{}

	response []interface{}

	iniElement InfoElement
	user       User
	Store      *Store
}

func (whendy *Whendy) Init() {

	str := tool.LoadFile(whendy.IniSource)

	json.Unmarshal([]byte(str), &whendy.iniElement)

}

func (w *Whendy) Element(e Element) {

	e.EvalMethod("nope")
}

func (wh *Whendy) EvalStartMode() {

	mode := wh.req.Header.Get("Application-Mode")

	if mode == "start" {
		fmt.Println(wh.req.Header.Get("Application-Mode"))

		wh.setElement(wh.iniElement)
	}

}
func (whendy *Whendy) Render() {

	whendy.EvalStartMode()
	w := whendy.w

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-API-KEY, Origin, X-Requested-With, Content-Type, Accept, Access-Control-Request-Method, Application-Mode, authorization, sid,  Application-Id")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	w.Header().Set("Allow", "GET, POST, OPTIONS, PUT, DELETE")

	w.Header().Set("Content-Type", "application/json")

	//fmt.Fprintf(w, `[9,4,8,16,32]`)

	jsonData, err := json.Marshal(whendy.getResponse())
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Fprint(w, string(jsonData))
	fmt.Println("x es ", X)

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

	whendy.user = User{}
	whendy.user.Init(w, req)

	whendy.w = w
	whendy.req = req
	whendy.Init()

	fmt.Println(whendy.Store.Session.SessionID())

	//whendy.Store.Session.Set("name", "yes")
	//fmt.Println("whendy.Store.GetSes -> ", whendy.Store.Session.Get("name"))

}

func (whendy *Whendy) setElement(info InfoElement) {

	typ, _ := element.New(info.Element)

	ele := typ.(Element)
	ele.SetStore(whendy.Store)
	ele.Init(info)
	ele.EvalMethod(info.Method)

	whendy.mergeResponse(ele.GetResponse())

	whendy.doUserAdmin(typ)
	whendy.doElementAdmin(typ)

	fmt.Println("user ", whendy.user.user)

}

func (whendy *Whendy) doUserAdmin(typ interface{}) {
	ele, ok := typ.(UserAdmin)

	if !ok {
		return
	}

	Print(ele.GetUserInfo())
}

func (whendy *Whendy) doElementAdmin(typ interface{}) {
	ele, ok := typ.(ElementAdmin)

	if !ok {
		return
	}

	Print(ele.GetElements())
}

func (whendy *Whendy) getResponse() []interface{} {

	return whendy.response
}

func (whendy *Whendy) mergeResponse(response []interface{}) {

	whendy.response = append(whendy.response, response...)

}

func Print(face interface{}) {
	jsonData, err := json.Marshal(face)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Printf("merge data Printing : %s\n", jsonData)
}
