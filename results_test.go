package go_iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterable_ToSlice(t *testing.T) {
	iter := New([]int{1, 2, 3, 4, 5})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, iter.ToSlice())
}

func TestIterable_ToChannel(t *testing.T) {
	iter := New([]int{1, 2, 3, 4, 5})

	channel := iter.ToChannel()

	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, <-channel)
	}
}

func TestIterable_ToMap(t *testing.T) {
	iter := New([]int{1, 2, 3, 4, 5})

	result := iter.ToMap()

	assert.Equal(t, map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}, result)
}
