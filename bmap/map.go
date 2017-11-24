package bmap

import "reflect"

func New(maplike interface{}) bmap {
	m := reflect.ValueOf(maplike)
	if m.Kind() != reflect.Map {
		panic("can not create a bmap from a not map type")
	}

	ret := make(bmap)

	for _, valueOfKey := range m.MapKeys() {
		ret[valueOfKey.Interface()] = m.MapIndex(valueOfKey).Interface()
	}

	return ret
}

type bmap map[interface{}]interface{}

//Values
//Keys
//Each
//Map
//Contains
