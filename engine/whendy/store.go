package whendy

import "sevian.com/whendy/tool"

// var session map[string]interface{}
var session = make(map[string]interface{})

type Store struct {
}

func (s *Store) SetSes(key string, value string) {

	session[key] = value
}

func (s *Store) GetSes(key string) interface{} {
	return session[key]

}

func (s *Store) LoadJson(path string) {
	data := tool.LoadJsonFile(path)

	for k, v := range data {
		//println(k, v)
		/*if _, ok := v.(string); ok {
			ss[k] = v
		}*/

		session[k] = v
	}

}
