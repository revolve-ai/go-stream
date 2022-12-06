package go_iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToSlice(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, iter.ToSlice())
}

func TestToChannel(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	channel := iter.ToChannel()

	for i := 1; i <= 5; i++ {
		assert.Equal(t, i, <-channel)
	}
}
