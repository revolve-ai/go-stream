package go_stream_test

import (
	"github.com/r-evolve/go-stream"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStream_Filter(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Filter(func(item int) bool {
		return item%2 == 0
	})

	assert.Equal(t, []int{2, 4, 6, 8, 10}, stream.ToSlice())
}

func TestStream_Map(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Map(func(item int) int {
		return item * 2
	})

	assert.Equal(t, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, stream.ToSlice())

	stream = go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestStream_Reduce(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.Reduce(func(acc, item int) int {
		return acc + item
	}, 0)

	assert.Equal(t, 55, result)
}

func TestStream_ForEach(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	counter := 0

	stream.ForEach(func(item int) {
		counter += 1
	})

	assert.Equal(t, 10, counter)
}

func TestStream_ForEachIndexed(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.ForEachIndexed(func(index int, item int) {
		_ = item * index
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, stream.ToSlice())
}

func TestStream_AnyMatch(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.AnyMatch(func(item int) bool {
		return item%2 == 0
	})

	assert.True(t, result)

	result = stream.AnyMatch(func(item int) bool {
		return item%11 == 0
	})

	assert.False(t, result)
}

func TestStream_AllMatch(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.AllMatch(func(item int) bool {
		return item%2 == 0
	})

	assert.False(t, result)

	result = stream.AllMatch(func(item int) bool {
		return item > 0
	})

	assert.True(t, result)
}

func TestStream_NoneMatch(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.NoneMatch(func(item int) bool {
		return item%2 == 0
	})

	assert.False(t, result)
}

func TestStream_FindFirst(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.FindFirst(func(item int) bool {
		return item%2 == 0
	})

	assert.Equal(t, 2, result)
}

func TestStream_Count(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.Count()

	assert.Equal(t, 10, result)
}

func TestStream_Min(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.Min(func(a, b int) int {
		return a - b
	})

	assert.Equal(t, 1, result)

	stream = go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result = stream.Min(func(a, b int) int {
		return b - a
	})

	assert.Equal(t, 10, result)
}

func TestStream_Max(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := stream.Max(func(a, b int) int {
		return a - b
	})

	assert.Equal(t, 10, result)
}

func TestStream_Limit(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Limit(5)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, stream.ToSlice())

	stream = go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Limit(15)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, stream.ToSlice())
}

func TestStream_Distinct(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Distinct()

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, stream.ToSlice())
}

func TestStream_Skip(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Skip(5)

	assert.Equal(t, []int{6, 7, 8, 9, 10}, stream.ToSlice())

	stream = go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Skip(15)

	assert.Equal(t, []int{}, stream.ToSlice())
}

func TestStream_Peek(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Peek(func(item int) {
		_ = item
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, stream.ToSlice())
}

func TestStream_Sort(t *testing.T) {
	stream := go_stream.New([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	stream.Sort(func(a, b int) int {
		return a - b
	})

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, stream.ToSlice())
}

func TestStream_Reverse(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	stream.Reverse()

	assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, stream.ToSlice())
}
