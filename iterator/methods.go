package iterator

func (iter *Iterator) Map(f func(interface{}) interface{}) *Iterator {
	iteratorCopy := NewIterator([]interface{}{})

	for item, _ := range iter.items {
		iteratorCopy.append(f(item))
	}

	return iteratorCopy
}

func (iter *Iterator) Filter(f func(interface{}) bool) *Iterator {
	iteratorCopy := NewIterator([]interface{}{})

	for item, _ := range iter.items {
		if f(item) {
			iteratorCopy.append(item)
		}
	}

	return iteratorCopy
}

func (iter *Iterator) Reduce(f func(interface{}, interface{}) interface{}) interface{} {
	var result interface{}

	for item, _ := range iter.items {
		result = f(result, item)
	}

	return result
}

func (iter *Iterator) First() interface{} {
	return iter.items[0]
}

func (iter *Iterator) Last() interface{} {
	return iter.items[len(iter.items)-1]
}

func (iter *Iterator) Count() int {
	return len(iter.items)
}

func (iter *Iterator) All(f func(interface{}) bool) bool {
	for item, _ := range iter.items {
		if !f(item) {
			return false
		}
	}

	return true
}

func (iter *Iterator) Any(f func(interface{}) bool) bool {
	for item, _ := range iter.items {
		if f(item) {
			return true
		}
	}

	return false
}

func (iter *Iterator) ForEach(f func(interface{})) {
	for item, _ := range iter.items {
		f(item)
	}
}

func (iter *Iterator) ToSlice() []interface{} {
	return iter.items
}

func (iter *Iterator) ToMap(f func(interface{}) (interface{}, interface{})) map[interface{}]interface{} {
	var m map[interface{}]interface{}

	for item, _ := range iter.items {
		key, value := f(item)

		m[key] = value
	}

	return m
}

func (iter *Iterator) ToSet() map[interface{}]bool {
	var set map[interface{}]bool

	for item, _ := range iter.items {
		set[item] = true
	}

	return set
}

func (iter *Iterator) ToChannel() chan interface{} {
	var channel chan interface{}

	for item, _ := range iter.items {
		channel <- item
	}

	return channel
}