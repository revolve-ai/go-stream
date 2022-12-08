package go_iterable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	stream := New([]int{1, 2, 3, 4, 5})

	assert.Equal(t, &Stream[int]{data: []int{1, 2, 3, 4, 5}}, stream)
}

func TestOf(t *testing.T) {
	stream := Of(1, 2, 3, 4, 5)

	assert.Equal(t, &Stream[int]{data: []int{1, 2, 3, 4, 5}}, stream)
}

func TestEmpty(t *testing.T) {
	stream := Empty[int]()

	assert.Equal(t, &Stream[int]{data: []int{}}, stream)
}

func TestStream_Append(t *testing.T) {
	stream := New([]int{1, 2, 3, 4, 5})

	stream.Append(6)

	assert.Equal(t, &Stream[int]{data: []int{1, 2, 3, 4, 5, 6}}, stream)
}

func TestStream_Remove(t *testing.T) {
	stream := New([]int{1, 2, 3, 4, 5})

	stream.Remove(3)

	assert.Equal(t, &Stream[int]{data: []int{1, 2, 4, 5}}, stream)
}
