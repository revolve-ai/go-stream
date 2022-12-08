package go_iterable

func (stream *Stream[T]) Filter(f Predicate[T]) *Stream[T] {
	for _, item := range stream.data {
		if !f(item) {
			stream.Remove(item)
		}
	}

	return stream
}

func (stream *Stream[T]) Map(f Mapper[T]) *Stream[T] {
	for index, item := range stream.data {
		stream.data[index] = f(item)
	}

	return stream
}

func (stream *Stream[T]) Reduce(f Reducer[T], initial T) T {
	result := initial

	for _, item := range stream.data {
		result = f(result, item)
	}

	return result
}

func (stream *Stream[T]) ForEach(f Consumer[T]) {
	for _, item := range stream.data {
		f(item)
	}
}

func (stream *Stream[T]) ForEachIndexed(f BiConsumer[int, T]) {
	for index, item := range stream.data {
		f(index, item)
	}
}

func (stream *Stream[T]) AnyMatch(f Predicate[T]) bool {
	for _, item := range stream.data {
		if f(item) {
			return true
		}
	}

	return false
}

func (stream *Stream[T]) AllMatch(f Predicate[T]) bool {
	for _, item := range stream.data {
		if !f(item) {
			return false
		}
	}

	return true
}

func (stream *Stream[T]) NoneMatch(f Predicate[T]) bool {
	return !stream.AnyMatch(f)
}

func (stream *Stream[T]) FindFirst(f Predicate[T]) T {
	return stream.Filter(f).data[0]
}

func (stream *Stream[T]) Count() int {
	return len(stream.data)
}

func (stream *Stream[T]) Min(f Comparator[T]) T {
	min := stream.data[0]

	for _, item := range stream.data {
		if f(item, min) < 0 {
			min = item
		}
	}

	return min
}

func (stream *Stream[T]) Max(f Comparator[T]) T {
	max := stream.data[0]

	for _, item := range stream.data {
		if f(item, max) > 0 {
			max = item
		}
	}

	return max
}

func (stream *Stream[T]) Limit(maxSize int) *Stream[T] {
	if maxSize > len(stream.data) {
		maxSize = len(stream.data)
	}

	stream.data = stream.data[:maxSize]

	return stream
}

func (stream *Stream[T]) Distinct() *Stream[T] {
	unique := make([]T, 0)

	for _, item := range stream.data {
		isUnique := true

		for _, uniqueItem := range unique {
			if item == uniqueItem {
				isUnique = false
			}
		}

		if isUnique {
			unique = append(unique, item)
		}
	}

	stream.data = unique

	return stream
}

func (stream *Stream[T]) Skip(n int) *Stream[T] {
	if n > len(stream.data) {
		n = len(stream.data)
	}

	stream.data = stream.data[n:]

	return stream
}

func (stream *Stream[T]) Peek(f Consumer[T]) *Stream[T] {
	for _, item := range stream.data {
		f(item)
	}

	return stream
}

func (stream *Stream[T]) Sort(f Comparator[T]) *Stream[T] {
	for i := range stream.data {
		for j := range stream.data {
			if f(stream.data[i], stream.data[j]) < 0 {
				stream.data[i], stream.data[j] = stream.data[j], stream.data[i]
			}
		}
	}

	return stream
}

func (stream *Stream[T]) Reverse() *Stream[T] {
	for i := 0; i < len(stream.data)/2; i++ {
		stream.data[i], stream.data[len(stream.data)-i-1] = stream.data[len(stream.data)-i-1], stream.data[i]
	}

	return stream
}
