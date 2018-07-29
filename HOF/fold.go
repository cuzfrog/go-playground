package HOF

import "reflect"

func foldLeft(col interface{}, id interface{}, fn func(prevId interface{}, next interface{}) interface{}) interface{} {
	c := reflect.ValueOf(col)
	for i := 0; i < c.Len(); i++ {
		id = fn(id, c.Index(i).Interface())
	}
	return id
}
