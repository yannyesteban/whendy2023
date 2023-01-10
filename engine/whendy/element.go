package whendy

type Element interface {
	EvalMethod()
	Init(map[string]interface{})
}
