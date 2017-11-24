package bmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
