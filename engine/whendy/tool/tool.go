package tool

import (
	"encoding/json"
	"os"
	"reflect"
)

func GetFileName(name string, pattern string) string {

	return "yanny"
}

func LoadJsonFile(name string) map[string]interface{} {

	dat, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	str := string(dat)

	byt := []byte(str)
	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	//print(data["BUILD"].(string))

	return data
}

func LoadArrayJsonFile(name string) []map[string]interface{} {

	dat, err := os.ReadFile(name)

	if err != nil {
		panic(err)
	}

	str := string(dat)

	byt := []byte(str)
	var data []map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	//print(data["BUILD"].(string))

	return data
}

func ConfigStructure(obj interface{}, data map[string]interface{}) {

	// pointer to struct - addressable
	ps := reflect.ValueOf(obj) //.Elem().FieldByName("Method1")

	// struct
	s := ps.Elem()

	if s.Kind() == reflect.Struct {
		for k, v := range data {
			field := s.FieldByName(k)
			if field.IsValid() {

				if field.CanSet() {
					// change value of Field

					if field.Kind() == reflect.Int {
						x := int64(v.(float64))
						field.SetInt(x)

					}

					if field.Kind() == reflect.String {
						field.SetString(v.(string))

					}
				}
			}
			//println(k, v.(string))
			/*if _, ok := v.(string); ok {
				ss[k] = v
			}*/

			//session[k] = v
		}

	}

}
