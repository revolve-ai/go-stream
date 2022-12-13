package go_stream

// A Predicate is a check on whether the given element fulfills a certain condition.
type Predicate[T comparable] func(T) bool

// A Comparator is a function which provides information on how to compare two values to each other.
// This may return -1 0 or 1 for where -1 says "is smaller", 0 says "is even" and 1 says "is greater".
type Comparator[T comparable] func(T, T) int

// A SimpleReducer is a function which combines a result value with a single item value.
// This is used to accumulate values e.g. []int{1,2,3,4,5} -> 15.
// WARNING: as this reducer is simple it might only work with same data types.
type SimpleReducer[T comparable] func(T, T) T

// A Reducer is a function which combines a result value with a single item value.
// This is used to unwrap data from a struct e.g. {value: 1} -> 1.
type Reducer[T comparable, R comparable] func(R, T) R

// A SimpleMapper is a function which is applied to each element in a Stream.
// This can be used to modify all values e.g. []int{1,2,3,4,5} -> []int{2,4,6,8,10} (here the mapper was x * 2)
// WARNING: as this mapper is simple it might only work with same data types.
type SimpleMapper[T comparable] func(T) T

// A Mapper is a function which is applied to each element in a Stream.
// This can be used to wrap all values e.g. []int{1,2} -> []SomeStruct{ {value: 1}, {value: 2} } (here the mapper was SomeStruct{value: x})
type Mapper[T comparable, R comparable] func(T) R

// A Consumer is a function which processes the given item independently of other items in a Stream.
type Consumer[T comparable] func(T)

// A BiConsumer is a function which processes two given items together (a pair) independently of other pairs in the Stream.
type BiConsumer[T comparable, U comparable] func(T, U)
