package bmap

import (
	"go-broom/bslice"
	"reflect"
)

func New() bmap {
	return NewOf(make(map[interface{}]interface{}))
}

func NewOf(maplike interface{}) bmap {
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

/*Utility*/
func (m bmap) Values() []interface{} {
	values := bslice.New()
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func (m bmap) Keys() []interface{} {
	keys := bslice.New()
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m bmap) Contains(searchKey interface{}) bool {
	_, found := m[searchKey]
	return found
}

/*Chainable*/
type iterator func(interface{}, interface{})

func (m bmap) Each(i iterator) bmap {
	for k, v := range m {
		i(k, v)
	}
	return m
}

type mapper func(interface{}, interface{}) interface{}

func (m bmap) Map(ma mapper) bmap {
	mapped := New()
	for i, v := range m {
		mapped[i] = ma(i, v)
	}
	return mapped
}
