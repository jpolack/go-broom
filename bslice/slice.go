package bslice

import (
	"reflect"
)

func New(slicelike interface{}) bslice {
	s := reflect.ValueOf(slicelike)
	if s.Kind() != reflect.Slice {
		panic("can not create a bslice from a not slice type")
	}

	ret := make(bslice, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

type bslice []interface{}

type iterator func(interface{}, int)

func (s bslice) Each(iter iterator) bslice {
	for i, v := range s {
		iter(v, i)
	}
	return s
}

type asnycIterator func(interface{}, int, chan<- bool)

func (s bslice) EachAsync(iter asnycIterator) bslice {
	c := make(chan bool, len(s))

	for i, v := range s {
		go iter(v, i, c)
	}

	for range s {
		<-c
	}
	return s
}

type mapper func(interface{}, int) interface{}

func (s bslice) Map(m mapper) bslice {
	mapped := make(bslice, len(s))
	for i, v := range s {
		mapped[i] = m(v, i)
	}
	return mapped
}

type reducer func(interface{}, interface{}, int) interface{}

func (s bslice) Reduce(r reducer, initialVal interface{}) interface{} {
	for i, v := range s {
		initialVal = r(initialVal, v, i)
	}
	return initialVal
}

type filter func(interface{}, int) bool

func (s bslice) Filter(f filter) bslice {
	filtered := bslice{}
	for i, v := range s {
		if f(v, i) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
