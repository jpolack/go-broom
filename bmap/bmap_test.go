package bmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	broomMap := New()

	assert.EqualValues(bmap{}, broomMap)
}
func TestNewSimple(t *testing.T) {
	assert := assert.New(t)

	original := map[int]int{1: 1, 2: 4, 3: 6}
	broomMap := New(original)
	broomMap[0] = 0

	assert.EqualValues(bmap{0: 0, 1: 1, 2: 4, 3: 6}, broomMap)
	assert.EqualValues(map[int]int{1: 1, 2: 4, 3: 6}, original)
}
func TestNewError(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		New("abc")
	})
}
func TestNewError2(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		New("abc", "def")
	})
}
func TestValues(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := New(normalMap)

	found := map[int]bool{2: false, 3: false, 4: false}

	for _, v := range broomMap.Values() {
		found[v.(int)] = true
	}

	assert.EqualValues(map[int]bool{2: true, 3: true, 4: true}, found)
}
func TestKeys(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := New(normalMap)

	found := map[int]bool{1: false, 2: false, 3: false}

	for _, v := range broomMap.Keys() {
		found[v.(int)] = true
	}

	assert.EqualValues(map[int]bool{1: true, 2: true, 3: true}, found)
}
func TestContains(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := New(normalMap)

	assert.True(broomMap.Contains(1))
	assert.False(broomMap.Contains(4))
}

func TestEach(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := New(normalMap)

	found := map[int]bool{1: false, 2: false, 3: false}

	broomMap.Each(func(key interface{}, value interface{}) {
		assert.EqualValues(normalMap[key.(int)], value.(int))
		found[key.(int)] = true
	})

	assert.EqualValues(map[int]bool{1: true, 2: true, 3: true}, found)
}

func TestMap(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := New(normalMap)

	found := map[int]bool{1: false, 2: false, 3: false}

	mapped := broomMap.Map(func(key interface{}, value interface{}) interface{} {
		assert.EqualValues(normalMap[key.(int)], value.(int))
		found[key.(int)] = true
		return value.(int) * 2
	})

	assert.EqualValues(map[int]bool{1: true, 2: true, 3: true}, found)
	assert.EqualValues(bmap{1: 4, 2: 6, 3: 8}, mapped)
}
