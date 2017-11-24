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
	broomMap := NewOf(original)
	broomMap[0] = 0

	assert.EqualValues(bmap{0: 0, 1: 1, 2: 4, 3: 6}, broomMap)
	assert.EqualValues(map[int]int{1: 1, 2: 4, 3: 6}, original)
}
func TestNewError(t *testing.T) {
	assert := assert.New(t)

	assert.Panics(func() {
		NewOf("abc")
	})
}
func TestValues(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := NewOf(normalMap)

	assert.EqualValues([]interface{}{2, 3, 4}, broomMap.Values())
}
func TestKeys(t *testing.T) {
	assert := assert.New(t)

	normalMap := map[int]int{1: 2, 2: 3, 3: 4}
	broomMap := NewOf(normalMap)

	assert.EqualValues([]interface{}{1, 2, 3}, broomMap.Keys())
}
