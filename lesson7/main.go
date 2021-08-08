package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

//Напишите функцию, которая на вход получает запрос SQL и произвольные параметры, среди которых могут быть как обычные значения (строка, число) так и слайсы таких значений.
//Позиция каждого переданного параметра в запросе SQL обозначается знаком "?".
//Функция должна вернуть запрос SQL, в котором для каждого параметра-слайса количество знаков "?" будет через запятую размножено до количества элементов слайса, а вторым ответом вернуть слайс из параметров, которые соответствуют новым позициям знаков "?".
//Пример:
//Вызов: func ( "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555 )
//Ответ: "SELECT * FROM table WHERE deleted = ? AND id IN(?,?,?) AND count < ?", []interface{}{ false, 1, 6, 234, 555 }
//
//Сделайте кодогенерацию с помощью easyjson для любой Вашей структуры.
//go:generate go run ./gen/main.go
//go:generate goimports -w ./assigns.go

func PrintStruct(q string, args interface{})(string, []interface{}) {
	if args == nil {
		panic("nill")
	}

	val := reflect.ValueOf(args)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		panic("non struct")
	}

	ret := make([]interface{}, 3)
	qs := ""

	for i := 0; i < val.NumField(); i++ {


		typeField := val.Type().Field(i)
		if typeField.Type.Kind() == reflect.Slice {
			qs = strings.Repeat("?", val.Field(i).Len())
			log.Printf("nested field: %v,%v", val.Field(i).Len())

			//q3 := strings.Index(q, "?")
			q = strings.Replace(q, "?", qs, val.Field(i).Len())



		}
		ret[i] = val.Field(i).Interface()
	}

	return q, ret


}

func main()  {
	m := struct{
		Bools bool
		SliceSql []int
		Ints int64
	}{
		Bools:  false,
		SliceSql: []int{1, 8, 234},
		Ints: 555,
	}
	q := "SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?"

	fmt.Println(PrintStruct(q, m))
}