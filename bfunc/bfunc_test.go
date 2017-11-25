package bfunc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOf(t *testing.T) {
	assert := assert.New(t)

	broomFunc := NewOf(func(x int, y int) int {
		return x + y
	})

	assert.EqualValues([]interface{}{3}, broomFunc(1, 2))
}
func TestApply(t *testing.T) {
	assert := assert.New(t)

	broomFunc := NewOf(func(x int, y int) int {
		return x + y
	})

	assert.EqualValues([]interface{}{3}, broomFunc.Args([]interface{}{1, 2}))
}
