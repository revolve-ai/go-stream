package go_iterable

func (iter *Iterable[T]) Map(f func(T) T) *Iterable[T] {
	iteratorCopy := NewIterable(make([]T, len(iter.items)))

	for _, item := range iter.items {
		iteratorCopy.append(f(item))
	}

	return iteratorCopy
}

func (iter *Iterable[T]) Filter(f func(T) bool) *Iterable[T] {
	iteratorCopy := NewIterable(make([]T, 0))

	for _, item := range iter.items {
		if f(item) {
			iteratorCopy.append(item)
		}
	}

	return iteratorCopy
}

func (iter *Iterable[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial

	for _, item := range iter.items {
		result = f(result, item)
	}

	return result
}

func (iter *Iterable[T]) First() T {
	return iter.items[0]
}

func (iter *Iterable[T]) Last() T {
	return iter.items[len(iter.items)-1]
}

func (iter *Iterable[T]) Count() int {
	return len(iter.items)
}

func (iter *Iterable[T]) All(f func(T) bool) bool {
	for _, item := range iter.items {
		if !f(item) {
			return false
		}
	}

	return true
}

func (iter *Iterable[T]) Any(f func(T) bool) bool {
	for _, item := range iter.items {
		if f(item) {
			return true
		}
	}

	return false
}

func (iter *Iterable[T]) ForEach(f func(T)) {
	for _, item := range iter.items {
		f(item)
	}
}

func (iter *Iterable[T]) ToSlice() []T {
	return iter.items
}

func (iter *Iterable[T]) ToChannel() chan T {
	channel := make(chan T, len(iter.items))

	for _, item := range iter.items {
		channel <- item
	}

	return channel
}
