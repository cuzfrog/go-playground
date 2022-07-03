package HOF

import "reflect"

func foldLeft(col interface{}, id interface{}, fn func(prevId interface{}, next interface{}) interface{}) interface{} {
	c := reflect.ValueOf(col)
	for i := 0; i < c.Len(); i++ {
		id = fn(id, c.Index(i).Interface())
	}
	return id
}

func foldLeftGeneric[T any, R any](col []T, id R, fn func(prev R, next T) R) R {
	acc := id
	for i := 0; i < len(col); i++ {
		acc = fn(acc, col[i])
	}
	return acc
}
