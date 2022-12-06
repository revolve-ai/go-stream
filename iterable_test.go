package go_iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIterable(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	assert.Equal(t, iter, &Iterable[int]{items: []int{1, 2, 3, 4, 5}, index: -1})
}

func TestIterable_Append(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.append(6)

	assert.Equal(t, iter, &Iterable[int]{items: []int{1, 2, 3, 4, 5, 6}, index: -1})
}

func TestIterable_Remove(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.remove(2)

	assert.Equal(t, iter, &Iterable[int]{items: []int{1, 2, 4, 5}, index: -1})
}

func TestIterable_Next(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	assert.True(t, iter.Next())
	assert.True(t, iter.Next())
	assert.True(t, iter.Next())
	assert.True(t, iter.Next())
	assert.True(t, iter.Next())
	assert.False(t, iter.Next())
}

func TestLoop(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	for iter.Next() {
		assert.Less(t, iter.Value(), 6)
	}
}

func TestIterable_Value(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.Next()
	iter.Next()
	iter.Next()
	iter.Next()
	iter.Next()

	assert.Equal(t, iter.Value(), 5)
}

func TestIterable_Reset(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.Next()
	iter.Next()
	iter.Next()
	iter.Next()
	iter.Next()

	iter.Reset()

	assert.Equal(t, iter.Value(), 1)
}
