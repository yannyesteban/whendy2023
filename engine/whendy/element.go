package whendy

type Element interface {
	//LoadConfig(string)
	EvalMethod(method string)
	Init(info InfoElement)
	GetResponse() []interface{}
	AddResponse(response interface{})
}
