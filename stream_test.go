package go_stream_test

import (
	go_stream "github.com/r-evolve/go-stream"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5})

	assert.Equal(t, []int{1, 2, 3, 4, 5}, stream.ToSlice())
}

func TestOf(t *testing.T) {
	stream := go_stream.Of(1, 2, 3, 4, 5)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, stream.ToSlice())
}

func TestEmpty(t *testing.T) {
	stream := go_stream.Empty[int]()

	assert.Equal(t, []int{}, stream.ToSlice())
}

func TestStream_Append(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5})

	stream.Append(6)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, stream.ToSlice())
}

func TestStream_Remove(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5})

	stream.Remove(3)

	assert.Equal(t, []int{1, 2, 4, 5}, stream.ToSlice())
}
