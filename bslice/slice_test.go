package bslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSimple(t *testing.T) {
	assert := assert.New(t)

	original := []int{1, 2, 3}
	broomSlice := New(original)
	broomSlice[1] = 3

	assert.EqualValues(bslice{1, 3, 3}, broomSlice)
	assert.EqualValues([]int{1, 2, 3}, original)
}
func TestNewError(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		New("abc")
	})
}

func TestEachSimple(t *testing.T) {
	assert := assert.New(t)

	iterable := New([]int{1, 2, 3})

	iterated := iterable.Each(func(v interface{}, i int) {
		iterable[i] = iterable[i].(int) + v.(int)
	})

	assert.EqualValues([]interface{}{2, 4, 6}, iterated)
}

func TestEachAsnycSimple(t *testing.T) {
	assert := assert.New(t)

	iterable := New([]int{1, 2, 3})

	newIterable := []int{}

	iterable.EachAsync(func(v interface{}, i int, c chan<- bool) {
		for j := 0; j < (10-i)*1000000; j++ {
		}
		newIterable = append(newIterable, v.(int))
		c <- true
	})

	assert.EqualValues([]int{3, 2, 1}, newIterable)
}

func TestMapSimple(t *testing.T) {
	assert := assert.New(t)

	mappable := New([]int{1, 2, 3})

	mapped := mappable.Map(func(v interface{}, i int) interface{} {
		return v.(int) + i
	})

	assert.EqualValues([]interface{}{1, 3, 5}, mapped)
}

func TestMapStruct(t *testing.T) {
	assert := assert.New(t)

	type test struct {
		a int
		b int
	}

	mappable := New([]test{
		test{1, 2},
		test{2, 3},
	})

	mapped := mappable.Map(func(v interface{}, i int) interface{} {
		testObject := v.(test)
		return testObject.a + testObject.b
	})

	assert.EqualValues([]interface{}{3, 5}, mapped)
}

func TestReduceSimple(t *testing.T) {
	assert := assert.New(t)

	reducable := New([]int{1, 2, 3})

	reduced := reducable.Reduce(func(sum interface{}, v interface{}, i int) interface{} {
		return sum.(int) + v.(int)
	}, 0)

	assert.EqualValues(6, reduced.(int))
}

func TestFilterSimple(t *testing.T) {
	assert := assert.New(t)

	filterable := New([]int{1, 2, 3})

	odds := filterable.Filter(func(v interface{}, i int) bool {
		return v.(int)%2 != 0
	})

	assert.EqualValues([]interface{}{1, 3}, odds)
}
