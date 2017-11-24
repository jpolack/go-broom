package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEachSimple(t *testing.T) {
	assert := assert.New(t)

	iterable := New([]int{1, 2, 3})

	iterated := iterable.Each(func(v interface{}, i int) {
		iterable[i] = iterable[i].(int) + v.(int)
	})

	assert.EqualValues([]interface{}{2, 4, 6}, iterated)
}
