package go_iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.Map(func(item int) int {
		return item * 2
	})

	assert.Equal(t, []int{2, 4, 6, 8, 10}, iter.items)
}

func TestFilter(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.Filter(func(item int) bool {
		return item%2 == 0
	})

	assert.Equal(t, []int{2, 4}, iter.items)
}

func TestReduce(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	result := iter.Reduce(func(result, item int) int {
		return result + item
	}, 0)

	assert.Equal(t, 15, result)
}

func TestFirst(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	assert.Equal(t, 1, iter.First())
}

func TestLast(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	assert.Equal(t, 5, iter.Last())
}

func TestCount(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	assert.Equal(t, 5, iter.Count())
}

func TestAll(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	result := iter.All(func(item int) bool {
		return item > 0
	})

	assert.True(t, result)
}

func TestAny(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	result := iter.Any(func(item int) bool {
		return item == 3
	})

	assert.True(t, result)
}

func TestForEach(t *testing.T) {
	iter := NewIterable([]int{1, 2, 3, 4, 5})

	iter.ForEach(func(item int) {
		item = item * 2
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, iter.items)
}

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
