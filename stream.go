package go_iterable

type Stream[T comparable] struct {
	data []T
}

func New[T comparable](data []T) *Stream[T] {
	stream := &Stream[T]{
		data: make([]T, 0),
	}

	for _, item := range data {
		stream.Append(item)
	}

	return stream
}

func Of[T comparable](data ...T) *Stream[T] {
	return New(data)
}

func Empty[T comparable]() *Stream[T] {
	return New([]T{})
}

func (stream *Stream[T]) Append(item T) {
	stream.data = append(stream.data, item)
}

func (stream *Stream[T]) Remove(item T) {
	for index, value := range stream.data {
		if value == item {
			stream.data = append(stream.data[:index], stream.data[index+1:]...)
		}
	}
}
