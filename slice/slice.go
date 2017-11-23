package slice

import (
	"reflect"
)

func New(slicelike interface{}) slice {
	s := reflect.ValueOf(slicelike)
	if s.Kind() != reflect.Slice {
		panic("toInterfaceSlice expects a slice")
	}

	ret := make(slice, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

type slice []interface{}

type iterator func(interface{}, int)

func (s slice) Each(iter iterator) slice {
	for i, v := range s {
		iter(v, i)
	}
	return s
}

type asnycIterator func(interface{}, int, chan<- bool)

func (s slice) EachAsync(iter asnycIterator) slice {
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

func (s slice) Map(m mapper) slice {
	mapped := make(slice, len(s))
	for i, v := range s {
		mapped[i] = m(v, i)
	}
	return mapped
}

type reducer func(interface{}, interface{}, int) interface{}

func (s slice) Reduce(r reducer, initialVal interface{}) interface{} {
	for i, v := range s {
		initialVal = r(initialVal, v, i)
	}
	return initialVal
}

type filter func(interface{}, int) bool

func (s slice) Filter(f filter) slice {
	filtered := slice{}
	for i, v := range s {
		if f(v, i) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}
