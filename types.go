package go_iterable

type Predicate[T comparable] func(T) bool
type Comparator[T comparable] func(T, T) int
type Reducer[T comparable] func(T, T) T
type Mapper[T comparable] func(T) T
type Consumer[T comparable] func(T)
type BiConsumer[T comparable, U comparable] func(T, U)
