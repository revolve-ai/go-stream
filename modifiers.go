package go_stream

// Filter is a method which can be used with any Stream to filter out unwanted values.
// To filter items correctly a Predicate is required.
// A filter is returning a filtered Stream to continue work.
// This method supports function chaining.
func (stream *Stream[T]) Filter(f Predicate[T]) *Stream[T] {
	for _, item := range stream.data {
		if !f(item) {
			stream.Remove(item)
		}
	}

	return stream
}

// Map is a method which can be used with any Stream to transform data.
// WARNING: this version of Map only works with the same Type as input and output!
// For Type-conversions see MapComplex.
// Map is returning a modified Stream to continue work.
// This method supports function chaining.
func (stream *Stream[T]) Map(f SimpleMapper[T]) *Stream[T] {
	for index, item := range stream.data {
		stream.data[index] = f(item)
	}

	return stream
}

// Reduce is a method which can be used with any Stream to transform data.
// WARNING: this version of Reduce only works with the same Type as input and output!
// For Type-conversions see ReduceComplex.
// Reduce is returning a modified Stream to continue work.
// This method supports function chaining.
func (stream *Stream[T]) Reduce(f SimpleReducer[T], initial T) T {
	result := initial

	for _, item := range stream.data {
		result = f(result, item)
	}

	return result
}

// ForEach is a method which can be used with any Stream to loop over each data point.
func (stream *Stream[T]) ForEach(f Consumer[T]) {
	for _, item := range stream.data {
		f(item)
	}
}

// ForEachIndexed is a method which can be used with any Stream to loop over each data point with an index.
func (stream *Stream[T]) ForEachIndexed(f BiConsumer[int, T]) {
	for index, item := range stream.data {
		f(index, item)
	}
}

// AnyMatch can be used with any Stream.
// This method is used to check whether any of the items inside the Stream matches the Predicate.
func (stream *Stream[T]) AnyMatch(f Predicate[T]) bool {
	for _, item := range stream.data {
		if f(item) {
			return true
		}
	}

	return false
}

// AllMatch can be used with any Stream.
// This method is used to check whether all the items inside the Stream matches the Predicate.
func (stream *Stream[T]) AllMatch(f Predicate[T]) bool {
	for _, item := range stream.data {
		if !f(item) {
			return false
		}
	}

	return true
}

// NoneMatch can be used with any Stream.
// This method is used to check whether none of the items inside the Stream matches the Predicate.
func (stream *Stream[T]) NoneMatch(f Predicate[T]) bool {
	return !stream.AnyMatch(f)
}

// FindFirst can be used with any Stream.
// This method is a shortcut to the Filter method.
// It is used to select the first item of the stream that matches the given Predicate.
func (stream *Stream[T]) FindFirst(f Predicate[T]) T {
	return stream.Filter(f).data[0]
}

// Count can be used with any Stream.
// This method counts the entries inside the given Stream instance.
func (stream *Stream[T]) Count() int {
	return len(stream.data)
}

// Min can be used with any Stream.
// This method selects the minimum value of a Stream instance based on the Comparator.
func (stream *Stream[T]) Min(f Comparator[T]) T {
	min := stream.data[0]

	for _, item := range stream.data {
		if f(item, min) < 0 {
			min = item
		}
	}

	return min
}

// Max can be used with any Stream.
// This method selects the maximum value of a Stream instance based on the Comparator.
func (stream *Stream[T]) Max(f Comparator[T]) T {
	max := stream.data[0]

	for _, item := range stream.data {
		if f(item, max) > 0 {
			max = item
		}
	}

	return max
}

// Limit can be used with any Stream.
// This method limits the size to the given number by cutting
// off all items with an index larger or equal than the given number.
func (stream *Stream[T]) Limit(maxSize int) *Stream[T] {
	if maxSize > len(stream.data) {
		maxSize = len(stream.data)
	}

	stream.data = stream.data[:maxSize]

	return stream
}

// Distinct can be used with any Stream.
// This method removes duplicate entries from the Stream.
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

// Skip can be used with any Stream.
// This method skips over the count n of items (removes the first n items) from the Stream.
func (stream *Stream[T]) Skip(n int) *Stream[T] {
	if n > len(stream.data) {
		n = len(stream.data)
	}

	stream.data = stream.data[n:]

	return stream
}

// Peek can be used with any Stream.
// This method works like a ForEach, but will return an unmodified Stream after looping.
func (stream *Stream[T]) Peek(f Consumer[T]) *Stream[T] {
	stream.ForEach(f)

	return stream
}

// Sort can be used with any Stream.
// This method sorts the Stream based on the given Comparator.
// WARNING: this is currently quite slow as bubble-sort is used.
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

// Reverse can be used with any Stream.
// This method reverses the items inside the Stream.
// The order will be exactly reversed afterwards.
func (stream *Stream[T]) Reverse() *Stream[T] {
	for i := 0; i < len(stream.data)/2; i++ {
		stream.data[i], stream.data[len(stream.data)-i-1] = stream.data[len(stream.data)-i-1], stream.data[i]
	}

	return stream
}
