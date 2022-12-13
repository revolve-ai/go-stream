package go_stream

// MapComplex is a more dynamic version of Map.
// This sadly cannot be implemented on the Stream struct itself, as there are no generics allowed.
func MapComplex[T comparable, R comparable](stream *Stream[T], f Mapper[T, R]) *Stream[R] {
	result := Empty[R]()

	for _, item := range stream.data {
		result.Append(f(item))
	}

	return result
}

// ReduceComplex is a more dynamic version of Reduce.
// This sadly cannot be implemented on the Stream struct itself, as there are no generics allowed.
func ReduceComplex[T comparable, R comparable](stream *Stream[T], f Reducer[T, R], initial R) *Stream[R] {
	result := initial

	for _, item := range stream.data {
		result = f(result, item)
	}

	return Of(result)
}
