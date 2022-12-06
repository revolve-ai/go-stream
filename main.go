package main

import (
	"fmt"
	"github.com/r-evolve/go-iterable/iterator"
)

func main() {
	iter := iterator.NewIterator([]int{1, 2, 3, 4, 5})

	fmt.Println(iter.Filter(func(item interface{}) bool {
		return item.(int) > 3
	}).First())

	fmt.Println(iter.Map(func(item interface{}) interface{} {
		return item.(int) * 2
	}))

	for iter.Next() {
		fmt.Println(iter.Value())
	}
}