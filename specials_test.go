package go_stream_test

import (
	"github.com/r-evolve/go-stream"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapComplex(t *testing.T) {
	stream := go_stream.New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	result := go_stream.MapComplex(stream, func(a int) float32 {
		return float32(a * 2)
	})

	assert.Equal(t, []float32{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}, result.ToSlice())
}

func TestReduceComplex(t *testing.T) {
	type ComplexType struct {
		value int
	}

	stream := go_stream.New([]ComplexType{
		{value: 1},
		{value: 2},
		{value: 3},
		{value: 4},
		{value: 5},
		{value: 6},
		{value: 7},
		{value: 8},
		{value: 9},
		{value: 10},
	})

	result := go_stream.ReduceComplex(stream, func(res int, a ComplexType) int {
		return res + a.value
	}, 0)

	assert.Equal(t, []int{55}, result.ToSlice())
}
