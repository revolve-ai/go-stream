package go_iterable

// ToSlice returns the data in the iterable as a slice.
func (stream *Stream[T]) ToSlice() []T {
	return stream.data
}

// ToChannel returns the data in the iterable as a channel.
func (stream *Stream[T]) ToChannel() chan T {
	channel := make(chan T, len(stream.data))

	for _, item := range stream.data {
		channel <- item
	}

	return channel
}

// ToMap returns the data in the iterable as a map.
func (stream *Stream[T]) ToMap() map[int]T {
	m := make(map[int]T, len(stream.data))

	for index, item := range stream.data {
		m[index] = item
	}

	return m
}
