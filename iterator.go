package go_iterable

type Iterator struct {
	items []interface{}
	index int
}

func NewIterator[T any](iterable []T) *Iterator {
	var iter Iterator

	for item := range iterable {
		iter.append(item)
	}

	return &iter
}

func (iter *Iterator) append(item interface{}) {
	iter.items = append(iter.items, item)
}

func (iter *Iterator) Next() bool {
	iter.index += 1

	return len(iter.items) > iter.index
}

func (iter *Iterator) Value() interface{} {
	return iter.items[iter.index]
}

func (iter *Iterator) Reset() {
	iter.index = 0
}

// TODO: requires fix
func (iter *Iterator) String() string {
	return iter.Reduce(func(result, item interface{}) interface{} {
		return result.(string) + item.(string)
	}).(string)
}
