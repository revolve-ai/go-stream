package go_iterable

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
