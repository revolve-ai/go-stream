package go_stream

// Stream is a type which works as a wrapper for slices.
// It enables to use common methods from the Java's Stream API
// to be used on Go's slices.
// A Stream is mutable by default, but will be able to be immutable in the future.
//
// TODO: implement mutable variant of stream
type Stream[T comparable] struct {
	data      []T
	isMutable bool
}

// New is the default constructor to create a Stream with a given slice.
func New[T comparable](data []T) *Stream[T] {
	return &Stream[T]{data, false}
}

// Of is a more dynamic constructor for Stream which provides the option to create a Stream of off single values.
func Of[T comparable](data ...T) *Stream[T] {
	return New(data)
}

// Empty will create an empty Stream.
func Empty[T comparable]() *Stream[T] {
	return New([]T{})
}

// Append works like Go's append function, just for a Stream.
func (stream *Stream[T]) Append(item T) {
	stream.data = append(stream.data, item)
}

// Remove will remove all occurrences of an item which is in the Stream.
func (stream *Stream[T]) Remove(item T) {
	for index, value := range stream.data {
		if value == item {
			stream.data = append(stream.data[:index], stream.data[index+1:]...)
		}
	}
}
