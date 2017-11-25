package bfunc

import (
	"go-broom/bslice"
	"reflect"
)

func NewOf(funclike interface{}) bfunc {
	f := reflect.ValueOf(funclike)
	if f.Kind() != reflect.Func {
		panic("can not create a bfunc from a not func type")
	}

	ret := bfunc(func(args ...interface{}) []interface{} {
		argsValues := make([]reflect.Value, len(args))
		bslice.NewOf(args).Each(func(v interface{}, i int) {
			argsValues[i] = reflect.ValueOf(v)
		})

		res := f.Call(argsValues)
		return bslice.NewOf(res).Map(func(v interface{}, i int) interface{} {
			return v.(reflect.Value).Interface()
		})
	})

	return ret
}

type bfunc func(...interface{}) []interface{}

func (f bfunc) Args(args []interface{}) []interface{} {
	return f(args...)
}
