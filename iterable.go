package go_iterable

type Iterable[T any] struct {
	items []T
	index int
}

func NewIterable[T any](iterable []T) *Iterable[T] {
	iter := &Iterable[T]{
		items: make([]T, 0),
		index: -1,
	}

	for _, item := range iterable {
		iter.append(item)
	}

	return iter
}

func (iter *Iterable[T]) append(item T) {
	iter.items = append(iter.items, item)
}

func (iter *Iterable[T]) remove(index int) {
	iter.items = append(iter.items[:index], iter.items[index+1:]...)
}

func (iter *Iterable[T]) Next() bool {
	if iter.index >= (len(iter.items) - 1) {
		return false
	}

	iter.index += 1

	return true
}

func (iter *Iterable[T]) Value() T {
	return iter.items[iter.index]
}

func (iter *Iterable[T]) Reset() {
	iter.index = 0
}
